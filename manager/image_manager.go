package manager

import (
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

func (img *Image) CreateImage(db *mgo.Database, user *User) {
	img.ID = bson.NewObjectId()
	img.UserID = user.ID

	if err := db.C("images").Insert(img); err != nil {
		panic(err)
	}
}

func (img *Image) ReadImageByID(db *mgo.Database) (found bool) {
	query := bson.M{"imageId": img.ImageID}

	if err := db.C("images").Find(query).One(img); err != nil {
		switch err {
		case mgo.ErrNotFound:
			return false
		default:
			panic(err)
		}
	}

	return true
}
