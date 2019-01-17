package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gchumillas/photomanager/handler"
	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, "mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}

	env := handler.NewEnv(client)

	r := mux.NewRouter()
	r.HandleFunc("/categories", env.GetCategories).Methods("GET")
	r.HandleFunc("/categories/{id}", env.GetCategory).Methods("GET")
	r.HandleFunc("/categories", env.PostCategory).Methods("POST")
	r.HandleFunc("/categories", env.PutCategory).Methods("PUT")

	// TODO: configuration needed
	log.Printf("Server started at port %v", "8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
