package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"net/textproto"
	"reflect"
	"strings"

	"github.com/gchumillas/photomanager/manager"
	"github.com/globalsign/mgo"
)

// Context key.
type contextKey string

type httpStatus struct {
	code int
	msg  string
}

var contextAuthUser = contextKey("auth-user")

// Common http errors.
var (
	payloadError        = httpStatus{400, "The payload is not well formed."}
	badParamsError      = httpStatus{400, "Bad parameters."}
	duplicateImageError = httpStatus{400, "Duplicate image."}
	invalidImageError   = httpStatus{400, "Invalid image format."}
	unauthorizedError   = httpStatus{401, "Not authorized."}
	docNotFoundError    = httpStatus{404, "Document not found."}
)

// Env contains common variables, such as the database access, etc.
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

	// Removes whitespaces around texts.
	elem := reflect.ValueOf(payload).Elem()
	switch reflect.TypeOf(elem).Kind() {
	case reflect.Struct:
		count := elem.NumField()
		for i := 0; i < count; i++ {
			field := elem.Field(i)
			if field.Type().Kind() == reflect.String {
				field.SetString(strings.Trim(field.String(), " "))
			}
		}
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

func getContentType(header textproto.MIMEHeader) string {
	cType := ""
	if len(header["Content-Type"]) > 0 {
		cType = header["Content-Type"][0]
	}

	return cType
}

func inArray(item string, items []string) bool {
	for index := range items {
		if items[index] == item {
			return true
		}
	}

	return false
}
