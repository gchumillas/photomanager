package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gchumillas/photomanager/handler"
	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var config struct {
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
	r.HandleFunc("/categories", env.GetCategories).Methods("GET")
	r.HandleFunc("/categories/{id}", env.GetCategory).Methods("GET")
	r.HandleFunc("/categories", env.PostCategory).Methods("POST")
	r.HandleFunc("/categories/{id}", env.PutCategory).Methods("PUT")

	// TODO: configuration needed
	log.Printf("Server started at port %v", config.ServerPort)
	log.Fatal(http.ListenAndServe(config.ServerPort, r))
}
