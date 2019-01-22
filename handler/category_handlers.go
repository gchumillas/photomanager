package handler

import (
	"net/http"

	"github.com/gchumillas/photomanager/manager"
)

func (env *Env) GetCategories(w http.ResponseWriter, r *http.Request) {
	manager.GetAllCategories(env.db)
	return
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
