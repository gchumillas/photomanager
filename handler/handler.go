package handler

import (
	"encoding/json"
	"net/http"

	"github.com/globalsign/mgo"
)

type httpStatus struct {
	msg  string
	code int
}

// Common http errors.
var (
	payloadError     httpStatus = httpStatus{"The payload is not well formed.", 400}
	docNotFoundError httpStatus = httpStatus{"The parameters are not valid.", 404}
	badParamsError   httpStatus = httpStatus{"Bad parameters.", 400}
)

type Env struct {
	db *mgo.Database
}

func NewEnv(db *mgo.Database) *Env {
	return &Env{db}
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
	return
}
