package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gchumillas/photomanager/dbox"
	"github.com/gchumillas/photomanager/manager"
)

// UploadImage uploads an image to the user's account.
func (env *Env) UploadImage(w http.ResponseWriter, r *http.Request, maxMemorySize int64) {
	u := getAuthUser(r)

	// TODO: check mimetype
	r.ParseMultipartForm(maxMemorySize)
	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	path := "/" + strings.TrimLeft(handler.Filename, "/")

	// TODO: if two images are identical, dropbox doesn't create a new image
	imageID := dbox.UploadFile(u.AccessToken, file, path)

	img := manager.NewImage()
	img.ImageID = imageID
	img.CreateImage(env.DB, u)

	json.NewEncoder(w).Encode(map[string]interface{}{"id": img.ID})
}
