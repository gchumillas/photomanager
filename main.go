package main

import (
	"log"
	"net/http"

	"github.com/gchumillas/photomanager/handler"
	"github.com/gorilla/mux"
)

func main() {
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// client, err := mongo.Connect(ctx, "mongodb://localhost:27017")

	r := mux.NewRouter()

	r.HandleFunc("/categories", handler.GetCategories).Methods("GET")
	r.HandleFunc("/categories/{id}", handler.GetCategory).Methods("GET")
	r.HandleFunc("/categories", handler.PostCategory).Methods("POST")
	r.HandleFunc("/categories", handler.PutCategory).Methods("PUT")

	// TODO: configuration needed
	log.Printf("Server started at port %v", "8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
