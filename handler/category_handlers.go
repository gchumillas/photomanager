package handler

import (
	"log"
	"net/http"

	"github.com/gchumillas/photomanager/manager"
)

func (env *Env) GetCategories(w http.ResponseWriter, r *http.Request) {
	var items []manager.Category
	if err := manager.GetCategories(env.db, &items); err != nil {
		log.Fatal(err)
	}

	for _, item := range items {
		log.Println(item.Name)
	}

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
