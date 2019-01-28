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
	APIVersion string
	ServerAddr string
	MongoURI   string
	MongoDB    string
	MongoUser  string
	MongoPass  string
}

func main() {
	conf, err := loadConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}

	session, err := mgo.Dial(conf.MongoURI)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	db := session.DB(conf.MongoDB)
	if err = db.Login(conf.MongoUser, conf.MongoPass); err != nil {
		log.Fatal(err)
	}

	// middlewares
	jsonCType := middleware.NewContentType("application/json; charset=utf-8")

	env := handler.NewEnv(db)
	prefix := fmt.Sprintf("/%s", strings.TrimLeft(conf.APIVersion, "/"))
	r := mux.NewRouter().StrictSlash(true)
	s := r.PathPrefix(prefix).Subrouter()

	// categories routes
	cats := s.PathPrefix("/categories").Subrouter()
	cats.HandleFunc("", env.GetCategories).Methods("GET")
	cats.HandleFunc("/{id}/subcategories", env.GetSubcategories).Methods("GET")
	cats.HandleFunc("/{id}", env.GetCategory).Methods("GET")
	cats.HandleFunc("", env.PostCategory).Methods("POST")
	cats.HandleFunc("/{id}", env.PutCategory).Methods("PUT")
	cats.Use(jsonCType.Middleware)

	log.Printf("Server started at port %v", conf.ServerAddr)
	log.Fatal(http.ListenAndServe(conf.ServerAddr, r))
}

func loadConfig(filename string) (conf config, err error) {
	conf = config{}

	file, err := os.Open(filename)
	if err != nil {
		return
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&conf)
	if err != nil {
		return
	}

	return
}
