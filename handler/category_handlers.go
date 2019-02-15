package handler

// TODO: remove bson library
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
	u := getAuthUser(r)

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
	options := manager.QueryOptions{
		Skip:  page * env.MaxItemsPerPage,
		Limit: env.MaxItemsPerPage,
		Sort:  sortCols,
	}
	u.GetCategories(env.DB, options, &items)

	// Gets the number of pages.
	numItems := u.GetNumCategories(env.DB, options)
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

// GetSubcategories prints all subcategories.
func (env *Env) GetSubcategories(w http.ResponseWriter, r *http.Request) {
	u := getAuthUser(r)
	params := mux.Vars(r)
	parentCatID := params["id"]

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
	options := manager.QueryOptions{
		Skip:  page * env.MaxItemsPerPage,
		Limit: env.MaxItemsPerPage,
		Sort:  sortCols,
	}
	u.GetSubcategories(env.DB, options, parentCatID, &items)

	// Gets the number of pages.
	numItems := u.GetNumSubcategories(env.DB, options, parentCatID)
	numPages := numItems / env.MaxItemsPerPage
	remainder := numItems % env.MaxItemsPerPage
	if remainder > 0 {
		numPages++
	}

	// TODO: change "page" by "numPages"
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

	// TODO: use manager.NewCategory(catID)
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

	cat := &manager.Category{
		ID: bson.ObjectIdHex(catID),
	}
	if found := cat.DeleteCategory(env.DB); !found {
		httpError(w, docNotFoundError)
		return
	}
}
