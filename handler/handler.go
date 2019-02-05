package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/globalsign/mgo"
)

type httpStatus struct {
	msg  string
	code int
}

// Common http errors.
var (
	payloadError     = httpStatus{"The payload is not well formed.", 400}
	docNotFoundError = httpStatus{"Document not found.", 404}
	badParamsError   = httpStatus{"Bad parameters.", 400}
)

// Env contains environment variables.
type Env struct {
	DB              *mgo.Database
	APIVersion      string
	MaxItemsPerPage int
	DropboxKey      string
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
