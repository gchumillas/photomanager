package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gchumillas/photomanager/manager"
	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
)

func (env *Env) GetCategories(w http.ResponseWriter, r *http.Request) {
	items := []manager.Category{}
	if err := manager.GetCategories(env.db, &items); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(items)
}

func (env *Env) GetSubcategories(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	categoryId := params["id"]

	if !bson.IsObjectIdHex(categoryId) {
		http.Error(w, "Bad ID", http.StatusBadRequest)
		return
	}

	items := []manager.Category{}
	if err := manager.GetSubcategories(env.db, categoryId, &items); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(items)
}

func (env *Env) GetCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	categoryId := params["id"]

	if !bson.IsObjectIdHex(categoryId) {
		http.Error(w, "The parameters are not valid.", http.StatusBadRequest)
		return
	}

	item := manager.Category{}
	if err := manager.GetCategory(env.db, categoryId, &item); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(item)
}

func (env *Env) InsertCategory(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Name string
	}

	d := json.NewDecoder(r.Body)
	if err := d.Decode(&payload); err != nil {
		http.Error(w, "The payload is not well formed.", http.StatusBadRequest)
		return
	}

	cat := &manager.Category{
		Name:     payload.Name,
		ImageIDs: []bson.ObjectId{},
	}
	manager.InsertCategory(env.db, cat)
}

func (env *Env) EditCategory(w http.ResponseWriter, r *http.Request) {
	// TODO: use a function
	params := mux.Vars(r)
	categoryId := params["id"]

	// TODO: use constants
	if !bson.IsObjectIdHex(categoryId) {
		http.Error(w, "The parameters are not valid.", http.StatusBadRequest)
		return
	}

	// TODO: use a function
	var payload struct {
		Name string
	}

	d := json.NewDecoder(r.Body)
	if err := d.Decode(&payload); err != nil {
		http.Error(w, "The payload is not well formed.", http.StatusBadRequest)
		return
	}

	cat := &manager.Category{
		ID:       bson.ObjectIdHex(categoryId),
		Name:     payload.Name,
		ImageIDs: []bson.ObjectId{},
	}
	manager.UpdateCategory(env.db, categoryId, cat)
}

func (env *Env) DeleteCategory(w http.ResponseWriter, r *http.Request) {
}
