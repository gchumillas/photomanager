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

func GetCategories(db *mgo.Database, items *[]Category) {
	if err := db.C("categories").Find(nil).All(items); err != nil {
		log.Fatal(err)
	}
}

func GetSubcategories(db *mgo.Database, categoryId string, items *[]Category) {
	filter := bson.M{"parentCategoryId": bson.ObjectIdHex(categoryId)}

	if err := db.C("categories").Find(filter).All(items); err != nil {
		log.Fatal(err)
	}
}

// TODO: this method shouldn't return any error
func GetCategory(db *mgo.Database, categoryId string, item *Category) error {
	return db.C("categories").FindId(bson.ObjectIdHex(categoryId)).One(item)
}

func InsertCategory(db *mgo.Database, item *Category) {
	if err := db.C("categories").Insert(item); err != nil {
		log.Fatal(err)
	}
}

func UpdateCategory(db *mgo.Database, categoryId string, item *Category) {
	// TODO: this line is too long
	if err := db.C("categories").UpdateId(bson.ObjectIdHex(categoryId), item); err != nil {
		log.Fatal(err)
	}
}
