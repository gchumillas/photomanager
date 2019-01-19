package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gchumillas/photomanager/handler"
	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var config struct {
	APIVersion string
	MongoURI   string
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, config.MongoURI)
	if err != nil {
		log.Fatal(err)
	}

	env := handler.NewEnv(client)
	r := mux.NewRouter()
	prefix := fmt.Sprintf("/%s", strings.TrimLeft(config.APIVersion, "/"))
	s := r.PathPrefix(prefix).Subrouter()
	s.HandleFunc("/categories", env.GetSubcategories).Methods("GET")
	s.HandleFunc("/subcategories/{categoryId}", env.GetSubcategories).Methods("GET")
	s.HandleFunc("/categories/{id}", env.GetCategory).Methods("GET")
	s.HandleFunc("/categories", env.PostCategory).Methods("POST")
	s.HandleFunc("/categories/{id}", env.PutCategory).Methods("PUT")

	// TODO: configuration needed
	log.Printf("Server started at port %v", config.ServerPort)
	log.Fatal(http.ListenAndServe(config.ServerPort, r))
}
