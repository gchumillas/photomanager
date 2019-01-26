package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gchumillas/photomanager/srvmanager"
)

func (env *Env) GetCategories(w http.ResponseWriter, r *http.Request) {
	var items []srvmanager.Category
	if err := srvmanager.GetCategories(env.db, &items); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(items)
}

func (env *Env) GetSubcategories(w http.ResponseWriter, r *http.Request) {
	return
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
