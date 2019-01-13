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
	log.Fatal(http.ListenAndServe(":8080", r))
}
