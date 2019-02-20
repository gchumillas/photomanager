package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gchumillas/photomanager/manager"
	"github.com/gorilla/mux"
)

// GetCategories gets all the categories.
func (env *Env) GetCategories(w http.ResponseWriter, r *http.Request) {
	u := getAuthUser(r)

	params := mux.Vars(r)
	parentCatID := params["catID"]

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

	options := manager.QueryOptions{
		Skip:  page * env.MaxItemsPerPage,
		Limit: env.MaxItemsPerPage,
		Sort:  sortCols,
	}
	items := u.GetCategories(env.DB, options, parentCatID)

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

// CreateCategory inserts a category.
func (env *Env) CreateCategory(w http.ResponseWriter, r *http.Request) {
	u := getAuthUser(r)

	var payload struct{ Name string }
	parsePayload(w, r, &payload)
	if len(payload.Name) == 0 {
		httpError(w, badParamsError)
		return
	}

	cat := manager.NewCategory()
	cat.Name = payload.Name
	cat.CreateCategory(env.DB, u)

	json.NewEncoder(w).Encode(map[string]interface{}{"id": cat.ID})
}

// ReadCategory gets a category.
func (env *Env) ReadCategory(w http.ResponseWriter, r *http.Request) {
	u := getAuthUser(r)

	params := mux.Vars(r)
	catID := params["catID"]

	cat := manager.NewCategory(catID)
	if !cat.ReadCategory(env.DB, u) {
		httpError(w, docNotFoundError)
		return
	}

	json.NewEncoder(w).Encode(cat)
}

// UpdateCategory updates a category.
func (env *Env) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	u := getAuthUser(r)

	params := mux.Vars(r)
	catID := params["catID"]

	var payload struct{ Name string }
	parsePayload(w, r, &payload)
	if len(payload.Name) == 0 {
		httpError(w, badParamsError)
		return
	}

	cat := manager.NewCategory(catID)
	cat.Name = payload.Name
	if !cat.UpdateCategory(env.DB, u) {
		httpError(w, docNotFoundError)
		return
	}
}

// DeleteCategory deletes a category.
func (env *Env) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	u := getAuthUser(r)

	params := mux.Vars(r)
	catID := params["catID"]

	cat := manager.NewCategory(catID)
	if !cat.DeleteCategory(env.DB, u) {
		httpError(w, docNotFoundError)
		return
	}
}

func (env *Env) AddImage(w http.ResponseWriter, r *http.Request) {
	u := getAuthUser(r)

	params := mux.Vars(r)
	catID := params["catID"]
	imgID := params["imgID"]

	cat := manager.NewCategory(catID)
	img := manager.NewImage(imgID)
	if !img.ReadImage(env.DB, u) || !cat.ReadCategory(env.DB, u) {
		httpError(w, docNotFoundError)
		return
	}

	if !cat.HasImage(env.DB, img) {
		cat.AddImage(env.DB, u, img)
	}
}

func (env *Env) RemoveImage(w http.ResponseWriter, r *http.Request) {
	u := getAuthUser(r)

	params := mux.Vars(r)
	catID := params["catID"]
	imgID := params["imgID"]

	cat := manager.NewCategory(catID)
	img := manager.NewImage(imgID)
	if !img.ReadImage(env.DB, u) || !cat.ReadCategory(env.DB, u) || !cat.HasImage(env.DB, img) {
		httpError(w, docNotFoundError)
		return
	}

	cat.RemoveImage(env.DB, u, img)
}
