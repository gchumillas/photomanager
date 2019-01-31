package handler

import (
	"encoding/json"
	"net/http"

	"github.com/globalsign/mgo"
)

// TODO: this can be variables
// Common http request errors.
const (
	httpPayloadError     = "The payload is not well formed."
	httpParamsError      = "The parameters are not valid."
	httpDocNotFoundError = "Document not found."
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
		http.Error(w, httpPayloadError, http.StatusBadRequest)
		return
	}
}
