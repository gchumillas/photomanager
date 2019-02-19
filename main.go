package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gchumillas/photomanager/handler"
	"github.com/globalsign/mgo"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type config struct {
	APIVersion         string `json:"apiVersion"`
	ServerAddr         string `json:"serverAddr"`
	MaxItemsPerPage    int    `json:"maxItemsPerPage"`
	MongoURI           string `json:"mongoUri"`
	MongoDB            string `json:"mongoDb"`
	MongoUser          string `json:"mongoUser"`
	MongoPass          string `json:"mongoPass"`
	DropboxAppKey      string `json:"dropboxAppKey"`
	DropboxAppSecret   string `json:"dropboxAppSecret"`
	DropboxRedirectURI string `json:"dropboxRedirectUri"`
}

func main() {
	// TODO: replace json by yaml or toml
	conf := loadConfig("config.json")

	session, err := mgo.Dial(conf.MongoURI)
	if err != nil {
		// TODO: replace log.Fatal by log.Panic
		log.Fatal(err)
	}
	defer session.Close()

	db := session.DB(conf.MongoDB)
	if err = db.Login(conf.MongoUser, conf.MongoPass); err != nil {
		log.Fatal(err)
	}

	env := &handler.Env{DB: db, MaxItemsPerPage: conf.MaxItemsPerPage}
	prefix := fmt.Sprintf("/%s", strings.TrimLeft(conf.APIVersion, "/"))
	r := mux.NewRouter()
	r.Use(env.JSONMiddleware)
	s := r.PathPrefix(prefix).Subrouter()

	// authentication (public routes)
	auth := s.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/url", func(w http.ResponseWriter, r *http.Request) {
		env.GetAuthURL(w, r, conf.DropboxAppKey, conf.DropboxRedirectURI)
	}).Methods("GET")
	auth.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		env.Login(w, r, conf.DropboxAppKey, conf.DropboxAppSecret, conf.DropboxRedirectURI)
	})

	// categories (private routes)
	cats := s.PathPrefix("/categories").Subrouter()
	cats.HandleFunc("", env.GetCategories).Methods("GET")
	cats.HandleFunc("", env.CreateCategory).Methods("POST")
	cats.HandleFunc("/{id}", env.ReadCategory).Methods("GET")
	cats.HandleFunc("/{id}/categories", env.GetCategories).Methods("GET")
	cats.HandleFunc("/{id}", env.UpdateCategory).Methods("PUT")
	cats.HandleFunc("/{id}", env.DeleteCategory).Methods("DELETE")
	cats.Use(env.AuthMiddleware)

	// images (private routes)
	imgs := s.PathPrefix("/images").Subrouter()
	imgs.HandleFunc("", env.UploadImage).Methods("POST")
	imgs.Use(env.AuthMiddleware)

	log.Printf("Server started at port %v", conf.ServerAddr)
	log.Fatal(http.ListenAndServe(conf.ServerAddr, handlers.RecoveryHandler()(r)))
}

func loadConfig(filename string) (conf config) {
	conf = config{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&conf)
	if err != nil {
		log.Fatal(err)
	}

	return
}
