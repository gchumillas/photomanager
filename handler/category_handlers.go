package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gchumillas/photomanager/manager"
	"github.com/gorilla/mux"
)

// GetCategories prints all categories.
func (env *Env) GetCategories(w http.ResponseWriter, r *http.Request) {
	u := getAuthUser(r)

	params := mux.Vars(r)
	parentCatID := params["id"]
	log.Println("parentCategoryId", parentCatID)

	sortCols := strings.Split(getParam(r, "sort", "name"), ",")
	for _, col := range sortCols {
		pos := strings.IndexRune(col, '-')
		str := col
		if pos > -1 {
			str = col[pos+1:]
		}

		if !inArray(str, []string{"name"}) {
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
	options := manager.QueryOptions{
		Skip:  page * env.MaxItemsPerPage,
		Limit: env.MaxItemsPerPage,
		Sort:  sortCols,
	}
	u.GetCategories(env.DB, options, parentCatID, &items)

	// Gets the number of pages.
	numItems := u.GetNumCategories(env.DB, options, parentCatID)
	numPages := numItems / env.MaxItemsPerPage
	remainder := numItems % env.MaxItemsPerPage
	if remainder > 0 {
		numPages++
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"items":    items,
		"numPages": numPages,
	})
}

// TODO: move this method under CreateCategory()
// ReadCategory prints a specific category.
func (env *Env) ReadCategory(w http.ResponseWriter, r *http.Request) {
	u := getAuthUser(r)

	params := mux.Vars(r)
	catID := params["id"]

	var cat *manager.Category
	if cat = u.ReadCategory(env.DB, catID); cat != nil {
		httpError(w, docNotFoundError)
		return
	}

	json.NewEncoder(w).Encode(cat)
}

// CreateCategory inserts a category.
func (env *Env) CreateCategory(w http.ResponseWriter, r *http.Request) {
	u := getAuthUser(r)

	var payload struct{ Name string }
	parsePayload(w, r, &payload)

	cat := &manager.Category{Name: payload.Name}
	id := u.CreateCategory(env.DB, cat)

	json.NewEncoder(w).Encode(map[string]interface{}{"id": id})
}

// UpdateCategory updates a category.
func (env *Env) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	u := getAuthUser(r)

	params := mux.Vars(r)
	catID := params["id"]

	// TODO: category name is required
	var payload struct{ Name string }
	parsePayload(w, r, &payload)

	cat := &manager.Category{Name: payload.Name}
	if !u.UpdateCategory(env.DB, catID, cat) {
		httpError(w, docNotFoundError)
		return
	}
}

// DeleteCategory deletes a category.
func (env *Env) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	u := getAuthUser(r)

	params := mux.Vars(r)
	catID := params["id"]

	if !u.DeleteCategory(env.DB, catID) {
		httpError(w, docNotFoundError)
		return
	}
}
