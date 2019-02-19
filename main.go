package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gchumillas/photomanager/handler"
	"github.com/globalsign/mgo"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/pelletier/go-toml"
)

type dropboxConfig struct {
	AppKey      string `toml:"appKey"`
	AppSecret   string `toml:"appSecret"`
	RedirectURI string `toml:"redirectUri"`
}

type mongoConfig struct {
	URI  string `toml:"uri"`
	DB   string `toml:"db"`
	User string `toml:"user"`
	Pass string `toml:"pass"`
}

type uploadsConfig struct {
	MaxMemorySize int64 `toml:"maxMemorySize"`
}

type config struct {
	APIVersion      string `toml:"apiVersion"`
	ServerAddr      string `toml:"serverAddr"`
	MaxItemsPerPage int    `toml:"maxItemsPerPage"`
	MaxUploadSize   int64  `toml:"maxUploadSize"`
	Mongo           mongoConfig
	Dropbox         dropboxConfig
	Uploads         uploadsConfig
}

func main() {
	conf := loadConfig("config.toml")

	session, err := mgo.Dial(conf.Mongo.URI)
	if err != nil {
		// TODO: replace log.Fatal by log.Panic
		log.Fatal(err)
	}
	defer session.Close()

	db := session.DB(conf.Mongo.DB)
	if err = db.Login(conf.Mongo.User, conf.Mongo.Pass); err != nil {
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
		env.GetAuthURL(w, r, conf.Dropbox.AppKey, conf.Dropbox.RedirectURI)
	}).Methods("GET")
	auth.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		env.Login(w, r, conf.Dropbox.AppKey, conf.Dropbox.AppSecret, conf.Dropbox.RedirectURI)
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
	imgs.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		env.UploadImage(w, r, conf.Uploads.MaxMemorySize)
	}).Methods("POST")
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

	decoder := toml.NewDecoder(file)
	err = decoder.Decode(&conf)
	if err != nil {
		log.Fatal(err)
	}

	return
}
