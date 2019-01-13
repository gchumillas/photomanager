package handler

import (
	"io"
	"net/http"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `{"alive": true}`)
}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	return
}
