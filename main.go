package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gchumillas/photomanager/handler"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting the server...")
	r := mux.NewRouter()
	r.HandleFunc("/home", handler.GetHome).Methods("GET")

	done := make(chan bool)
	go http.ListenAndServe(":8080", r)
	log.Printf("Server started at port %v", "8080")
	<-done
}
