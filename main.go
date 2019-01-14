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
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, "mongodb://localhost:27017")

	r := mux.NewRouter()
	r.HandleFunc("/home", handler.GetHome).Methods("GET")

	done := make(chan bool)
	go http.ListenAndServe(":8080", r)
	log.Printf("Server started at port %v", "8080")
	<-done
}
