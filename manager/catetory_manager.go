package manager

import (
	"log"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type Category struct {
	ID       bson.ObjectId   `json:"id" bson:"_id,omitempty"`
	Name     string          `json:"name"`
	ImageIDs []bson.ObjectId `json:"imageIds" bson:"imageIds"`
}

// TODO: pagination
func GetCategories(db *mgo.Database, items *[]Category) {
	if err := db.C("categories").Find(nil).All(items); err != nil {
		log.Fatal(err)
	}
}

// TODO: pagination
func GetSubcategories(db *mgo.Database, catId string, items *[]Category) {
	filter := bson.M{"parentCategoryId": bson.ObjectIdHex(catId)}

	if err := db.C("categories").Find(filter).All(items); err != nil {
		log.Fatal(err)
	}
}

func GetCategory(db *mgo.Database, catId string, item *Category) (found bool) {
	id := bson.ObjectIdHex(catId)

	if err := db.C("categories").FindId(id).One(item); err != nil {
		switch err {
		case mgo.ErrNotFound:
			return false
		default:
			log.Fatal(err)
		}
	}

	return true
}

func InsertCategory(db *mgo.Database, item *Category) {
	if err := db.C("categories").Insert(item); err != nil {
		log.Fatal(err)
	}
}

func UpdateCategory(db *mgo.Database, catId string, item *Category) (found bool) {
	id := bson.ObjectIdHex(catId)

	if err := db.C("categories").UpdateId(id, item); err != nil {
		switch err {
		case mgo.ErrNotFound:
			return false
		default:
			log.Fatal(err)
		}
	}

	return true
}
