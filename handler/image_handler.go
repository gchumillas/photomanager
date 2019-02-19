package handler

import (
	"log"
	"net/http"

	"github.com/gchumillas/photomanager/dbox"
)

func (env *Env) UploadImage(w http.ResponseWriter, r *http.Request) {
	u := getAuthUser(r)

	// TODO: 32 Mebibits must be in config
	r.ParseMultipartForm(32 << 10)
	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	dbox.UploadFile(u.AccessToken, file, handler.Filename)
}
