package handler

import (
	"encoding/json"
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
		panic(err)
	}
	defer file.Close()

	path := "/" + strings.TrimLeft(handler.Filename, "/")
	imageID := dbox.UploadFile(u.AccessToken, file, path)

	img := manager.NewImage()
	img.ImageID = imageID
	if img.ReadImageByID(env.DB) {
		httpError(w, duplicateImageError)
		return
	}
	img.CreateImage(env.DB, u)

	json.NewEncoder(w).Encode(map[string]interface{}{"id": img.ID})
}
