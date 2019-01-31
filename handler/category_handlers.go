package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gchumillas/photomanager/manager"
	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
)

func (env *Env) GetCategories(w http.ResponseWriter, r *http.Request) {
	items := []manager.Category{}
	manager.GetCategories(env.db, &items)

	json.NewEncoder(w).Encode(items)
}

func (env *Env) GetSubcategories(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	categoryId := params["id"]

	// TODO: verify that category exists
	if !bson.IsObjectIdHex(categoryId) {
		http.Error(w, "Bad ID", http.StatusBadRequest)
		return
	}

	items := []manager.Category{}
	manager.GetSubcategories(env.db, categoryId, &items)

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
		http.Error(w, "Document not found.", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(item)
}

func (env *Env) InsertCategory(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Name string
	}
	parsePayload(w, r, &payload)

	cat := &manager.Category{
		Name:     payload.Name,
		ImageIDs: []bson.ObjectId{},
	}
	manager.InsertCategory(env.db, cat)
}

func (env *Env) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	categoryId := params["id"]

	if !bson.IsObjectIdHex(categoryId) {
		http.Error(w, badParamsErrror, http.StatusBadRequest)
		return
	}

	var payload struct {
		Name string
	}
	parsePayload(w, r, &payload)

	cat := &manager.Category{
		ID:       bson.ObjectIdHex(categoryId),
		Name:     payload.Name,
		ImageIDs: []bson.ObjectId{},
	}
	manager.UpdateCategory(env.db, categoryId, cat)
}

func (env *Env) DeleteCategory(w http.ResponseWriter, r *http.Request) {
}
