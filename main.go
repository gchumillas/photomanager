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

	// TODO: avoid using context. It adds unnecessary complexity to the code. Consider using another approach, such as a class wrapper.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, config.MongoURI)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(config.MongoDB)
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
