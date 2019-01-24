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

	db, err := dbConnect(conf)
	if err != nil {
		log.Fatal(err)
	}

	env := handler.NewEnv(db)
	prefix := fmt.Sprintf("/%s", strings.TrimLeft(conf.APIVersion, "/"))
	r := mux.NewRouter()
	s := r.PathPrefix(prefix).Subrouter()

	// categories routes
	s.HandleFunc("/categories", env.GetSubcategories).Methods("GET")
	s.HandleFunc("/subcategories/{categoryId}", env.GetSubcategories).Methods("GET")
	s.HandleFunc("/categories/{id}", env.GetCategory).Methods("GET")
	s.HandleFunc("/categories", env.PostCategory).Methods("POST")
	s.HandleFunc("/categories/{id}", env.PutCategory).Methods("PUT")

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

func dbConnect(conf config) (db *mgo.Database, err error) {
	session, err := mgo.Dial(conf.MongoURI)
	if err != nil {
		return
	}
	defer session.Close()

	err = session.Ping()
	if err != nil {
		return
	}

	db = session.DB(conf.MongoDB)
	if err = db.Login(conf.MongoUser, conf.MongoPass); err != nil {
		return
	}

	return
}
