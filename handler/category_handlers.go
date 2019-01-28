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
		http.Error(w, "Bad ID", http.StatusBadRequest)
		return
	}

	item := manager.Category{}
	if err := manager.GetCategory(env.db, categoryId, &item); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(item)
}

func (env *Env) PostCategory(w http.ResponseWriter, r *http.Request) {
	return
}

func (env *Env) PutCategory(w http.ResponseWriter, r *http.Request) {
	return
}
