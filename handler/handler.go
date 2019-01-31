package handler

import (
	"encoding/json"
	"net/http"

	"github.com/globalsign/mgo"
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
		http.Error(w, "The payload is not well formed.", http.StatusBadRequest)
		return
	}
}
