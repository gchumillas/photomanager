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
// TODO: categories should belongs to a specific user
func (env *Env) GetCategories(w http.ResponseWriter, r *http.Request) {
	parentCatID := getParam(r, "parentCatId", "")
	var query interface{}
	if len(parentCatID) > 0 {
		if !bson.IsObjectIdHex(parentCatID) {
			httpError(w, badParamsError)
			return
		}

		query = bson.M{"parentCategoryId": bson.ObjectIdHex(parentCatID)}
	}

	sortCols := strings.Split(getParam(r, "sort", "name"), ",")
	for _, col := range sortCols {
		str := col
		if i := strings.IndexRune(col, '-'); i == 0 {
			str = col[1:]
		}

		if found, _ := inArray(str, []string{"name"}); !found {
			httpError(w, badParamsError)
			return
		}
	}

	page, err := strconv.Atoi(getParam(r, "page", "0"))
	if err != nil || page < 0 {
		httpError(w, badParamsError)
		return
	}

	items := []manager.Category{}
	filter := manager.Filter{
		Skip:     page * env.MaxItemsPerPage,
		Limit:    env.MaxItemsPerPage,
		Query:    query,
		SortCols: sortCols,
	}
	manager.GetCategories(env.DB, filter, &items)

	// Gets the number of pages.
	numItems := manager.GetNumCategories(env.DB, filter)
	numPages := numItems / env.MaxItemsPerPage
	remainder := numItems % env.MaxItemsPerPage
	if remainder > 0 {
		numPages++
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"items": items,
		"page": map[string]interface{}{
			"current":  page,
			"total":    numPages,
			"maxItems": env.MaxItemsPerPage,
		},
	})
}

// ReadCategory prints a specific category.
func (env *Env) ReadCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	catID := params["id"]

	if !bson.IsObjectIdHex(catID) {
		httpError(w, badParamsError)
		return
	}

	cat := &manager.Category{ID: bson.ObjectIdHex(catID)}
	if found := cat.ReadCategory(env.DB); !found {
		httpError(w, docNotFoundError)
		return
	}

	json.NewEncoder(w).Encode(cat)
}

// CreateCategory inserts a category.
func (env *Env) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Name string
	}
	parsePayload(w, r, &payload)

	cat := &manager.Category{
		ID:   bson.NewObjectId(),
		Name: payload.Name,
	}
	cat.CreateCategory(env.DB)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"id": cat.ID,
	})
}

// UpdateCategory updates a category.
func (env *Env) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	catID := params["id"]

	if !bson.IsObjectIdHex(catID) {
		httpError(w, badParamsError)
		return
	}

	var payload struct {
		Name string
	}
	parsePayload(w, r, &payload)

	cat := &manager.Category{
		ID:       bson.ObjectIdHex(catID),
		Name:     payload.Name,
		ImageIDs: []bson.ObjectId{},
	}
	if found := cat.UpdateCategory(env.DB); !found {
		httpError(w, docNotFoundError)
		return
	}
}

// DeleteCategory deletes a category.
func (env *Env) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	catID := params["id"]

	if !bson.IsObjectIdHex(catID) {
		httpError(w, badParamsError)
		return
	}

	cat := &manager.Category{
		ID: bson.ObjectIdHex(catID),
	}
	if found := cat.DeleteCategory(env.DB); !found {
		httpError(w, docNotFoundError)
		return
	}
}
