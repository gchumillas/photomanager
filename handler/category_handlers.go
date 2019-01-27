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
	var items []manager.Category
	if err := manager.GetCategories(env.db, &items); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(items)
}

func (env *Env) GetSubcategories(w http.ResponseWriter, r *http.Request) {
	var items []manager.Category

	params := mux.Vars(r)
	categoryId := params["id"]

	if !bson.IsObjectIdHex(categoryId) {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	if err := manager.GetSubcategories(env.db, categoryId, &items); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(items)
}

func (env *Env) GetCategory(w http.ResponseWriter, r *http.Request) {
	return
}

func (env *Env) PostCategory(w http.ResponseWriter, r *http.Request) {
	return
}

func (env *Env) PutCategory(w http.ResponseWriter, r *http.Request) {
	return
}
