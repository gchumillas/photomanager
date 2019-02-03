package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gchumillas/photomanager/manager"
	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
)

// GetCategories prints all categories.
func (env *Env) GetCategories(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	parentCatID := params["id"]

	var query interface{}
	if len(parentCatID) > 0 {
		if !bson.IsObjectIdHex(parentCatID) {
			httpError(w, badParamsError)
			return
		}

		query = bson.M{"parentCategoryId": bson.ObjectIdHex(parentCatID)}
	}

	// TODO: check columns
	sortCols := strings.Split(getParam(r, "sort", "name"), ",")

	// page := strconv.Atoi(getParam("page", "0"))
	page, err := strconv.Atoi(getParam(r, "page", "0"))
	if err != nil || page < 0 {
		httpError(w, badParamsError)
		return
	}

	items := []manager.Category{}
	filter := manager.Filter{
		Skip:     page * env.maxItemsPerPage,
		Limit:    env.maxItemsPerPage,
		Query:    query,
		SortCols: sortCols,
	}
	manager.GetCategories(env.db, filter, &items)

	// Gets the number of pages.
	numItems := manager.GetNumCategories(env.db, filter)
	numPages := numItems / env.maxItemsPerPage
	remainder := numItems % env.maxItemsPerPage
	if remainder > 0 {
		numPages++
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"items": items,
		"page": map[string]interface{}{
			"current":  page,
			"total":    numPages,
			"maxItems": env.maxItemsPerPage,
		},
	})
}

// GetCategory prints a specific category.
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

// InsertCategory inserts a category.
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

// UpdateCategory updates a category.
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

// DeleteCategory deletes a category.
func (env *Env) DeleteCategory(w http.ResponseWriter, r *http.Request) {
}
