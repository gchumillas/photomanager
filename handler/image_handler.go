package handler

import (
	"log"
	"net/http"

	"github.com/gchumillas/photomanager/dbox"
)

func (env *Env) UploadImage(w http.ResponseWriter, r *http.Request, maxMemorySize int64) {
	u := getAuthUser(r)

	r.ParseMultipartForm(maxMemorySize)
	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	dbox.UploadFile(u.AccessToken, file, handler.Filename)
}
