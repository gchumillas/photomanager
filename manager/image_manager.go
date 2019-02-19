package manager

import (
	"log"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type Image struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	UserID  bson.ObjectId `json:"userId" bson:"userId"`
	ImageID string        `json:"imageId" bson:"imageId"`
}

// NewImage returns a new image.
func NewImage(catID ...string) *Image {
	var id bson.ObjectId
	if len(catID) > 0 {
		id = bson.ObjectIdHex(catID[0])
	}

	return &Image{ID: id}
}

func (image *Image) CreateImage(db *mgo.Database, user *User) {
	image.ID = bson.NewObjectId()
	image.UserID = user.ID

	if err := db.C("images").Insert(image); err != nil {
		log.Panic(err)
	}
}
