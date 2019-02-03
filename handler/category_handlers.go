package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gchumillas/photomanager/manager"
	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
)

// GetCategories gets all categories.
func (env *Env) GetCategories(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	parentCatID := params["id"]
	colsParam := params["columns"]

	var query interface{}
	if len(parentCatID) > 0 {
		if !bson.IsObjectIdHex(parentCatID) {
			httpError(w, badParamsError)
			return
		}

		query = bson.M{"parentCategoryId": bson.ObjectIdHex(parentCatID)}
	}

	var sortCols []string
	if len(colsParam) > 0 {
		// TODO: check columns
		sortCols = strings.Split(colsParam, ",")
	}

	// page := strconv.Atoi(getParam("page", "0"))

	pageParam := r.FormValue("page")
	page := 0
	if len(pageParam) > 0 {
		var err error
		page, err = strconv.Atoi(pageParam)
		if err != nil {
			httpError(w, badParamsError)
			return
		}

		log.Println(page)
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
