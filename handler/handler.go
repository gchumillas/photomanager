package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gchumillas/photomanager/manager"
	"github.com/globalsign/mgo"
)

// Context key.
type contextKey string

type httpStatus struct {
	msg  string
	code int
}

var contextAuthUser = contextKey("auth-user")

// Common http errors.
var (
	payloadError      = httpStatus{"The payload is not well formed.", 400}
	docNotFoundError  = httpStatus{"Document not found.", 404}
	badParamsError    = httpStatus{"Bad parameters.", 400}
	unauthorizedError = httpStatus{"Not authorized.", 401}
)

// Env contains environment variables.
type Env struct {
	DB              *mgo.Database
	MaxItemsPerPage int
}

func getAuthUser(r *http.Request) *manager.User {
	return r.Context().Value(contextAuthUser).(*manager.User)
}

func parsePayload(w http.ResponseWriter, r *http.Request, payload interface{}) {
	dec := json.NewDecoder(r.Body)

	if err := dec.Decode(payload); err != nil {
		httpError(w, payloadError)
		return
	}
}

func httpError(w http.ResponseWriter, status httpStatus) {
	http.Error(w, status.msg, status.code)
	log.Printf("http error (%d): %s", status.code, status.msg)
	return
}

func getParam(r *http.Request, key, def string) (param string) {
	if param = r.FormValue(key); len(param) == 0 {
		param = def
	}

	return
}

func inArray(item string, items []string) (found bool, index int) {
	for index = range items {
		if found = (items[index] == item); found {
			return
		}
	}

	return
}

// TODO: remove this function
func colsInArray(cols []string, items ...string) bool {
	for _, col := range cols {
		str := col
		if i := strings.IndexRune(col, '-'); i == 0 {
			str = col[1:]
		}

		if found, _ := inArray(str, items); !found {
			return false
		}
	}

	return true
}
