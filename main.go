package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gchumillas/photomanager/handler"
	"github.com/gchumillas/photomanager/middleware"
	"github.com/globalsign/mgo"
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
	DropboxAuthURL     string `json:"dropboxAuthUrl"`
	DropboxAppKey      string `json:"dropboxAppKey"`
	DropboxRedirectURI string `json:"dropboxRedirectUri"`
}

func main() {
	conf := loadConfig("config.json")

	session, err := mgo.Dial(conf.MongoURI)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	db := session.DB(conf.MongoDB)
	if err = db.Login(conf.MongoUser, conf.MongoPass); err != nil {
		log.Fatal(err)
	}

	env := &handler.Env{
		DB:              db,
		MaxItemsPerPage: conf.MaxItemsPerPage,
	}
	prefix := fmt.Sprintf("/%s", strings.TrimLeft(conf.APIVersion, "/"))
	r := mux.NewRouter()
	s := r.PathPrefix(prefix).Subrouter()

	// categories routes
	cats := s.PathPrefix("/categories").Subrouter()
	cats.HandleFunc("", env.GetCategories).Methods("GET")
	cats.HandleFunc("", env.CreateCategory).Methods("POST")
	cats.HandleFunc("/{id}", env.ReadCategory).Methods("GET")
	cats.HandleFunc("/{id}", env.UpdateCategory).Methods("PUT")
	cats.HandleFunc("/{id}", env.DeleteCategory).Methods("DELETE")

	// authentication
	auth := s.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/url", func(w http.ResponseWriter, r *http.Request) {
		env.GetAuthURL(
			w, r, conf.DropboxAuthURL, conf.DropboxRedirectURI, conf.DropboxAppKey,
		)
	}).Methods("GET")

	// middlewares
	cType := middleware.NewContentType("application/json; charset=utf-8")
	r.Use(cType.Middleware)

	log.Printf("Server started at port %v", conf.ServerAddr)
	log.Fatal(http.ListenAndServe(conf.ServerAddr, r))
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
