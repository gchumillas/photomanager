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

var config struct {
	APIVersion string
	MongoURI   string
	MongoDB    string
	ServerPort string
}

func main() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(file)
	error := decoder.Decode(&config)
	if error != nil {
		log.Fatal(error)
	}

	// Connects to MongoDB.
	session, err := mgo.Dial(config.MongoURI)
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Checks connection.
	err = session.Ping()
	if err != nil {
		log.Fatal(err)
	}

	db := session.DB(config.MongoDB)
	env := handler.NewEnv(db)

	prefix := fmt.Sprintf("/%s", strings.TrimLeft(config.APIVersion, "/"))
	r := mux.NewRouter()
	s := r.PathPrefix(prefix).Subrouter()

	// categories routes
	s.HandleFunc("/categories", env.GetSubcategories).Methods("GET")
	s.HandleFunc("/subcategories/{categoryId}", env.GetSubcategories).Methods("GET")
	s.HandleFunc("/categories/{id}", env.GetCategory).Methods("GET")
	s.HandleFunc("/categories", env.PostCategory).Methods("POST")
	s.HandleFunc("/categories/{id}", env.PutCategory).Methods("PUT")

	log.Printf("Server started at port %v", config.ServerPort)
	log.Fatal(http.ListenAndServe(config.ServerPort, r))
}
