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
	catId := params["id"]

	if !bson.IsObjectIdHex(catId) {
		httpError(w, badParamsError)
		return
	}

	items := []manager.Category{}
	manager.GetSubcategories(env.db, catId, &items)

	json.NewEncoder(w).Encode(items)
}

func (env *Env) GetCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	catId := params["id"]

	if !bson.IsObjectIdHex(catId) {
		httpError(w, badParamsError)
		return
	}

	item := manager.Category{}
	if found := manager.GetCategory(env.db, catId, &item); !found {
		httpError(w, docNotFoundError)
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
	catId := params["id"]

	if !bson.IsObjectIdHex(catId) {
		httpError(w, badParamsError)
		return
	}

	var payload struct {
		Name string
	}
	parsePayload(w, r, &payload)

	cat := &manager.Category{
		ID:       bson.ObjectIdHex(catId),
		Name:     payload.Name,
		ImageIDs: []bson.ObjectId{},
	}
	if found := manager.UpdateCategory(env.db, catId, cat); !found {
		httpError(w, docNotFoundError)
		return
	}
}

func (env *Env) DeleteCategory(w http.ResponseWriter, r *http.Request) {
}
