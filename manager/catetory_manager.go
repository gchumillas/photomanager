package manager

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type Category struct {
	ID       bson.ObjectId   `json:"id" bson:"_id,omitempty"`
	Name     string          `json:"name"`
	ImageIDs []bson.ObjectId `json:"imageIds" bson:"imageIds"`
}

func GetCategories(db *mgo.Database, items *[]Category) error {
	return db.C("categories").Find(nil).All(items)
}

func GetSubcategories(db *mgo.Database, categoryId string, items *[]Category) error {
	filter := bson.M{"parentCategoryId": bson.ObjectIdHex(categoryId)}

	return db.C("categories").Find(filter).All(items)
}

func GetCategory(db *mgo.Database, categoryId string, item *Category) error {
	return db.C("categories").FindId(bson.ObjectIdHex(categoryId)).One(item)
}

func InsertCategory(db *mgo.Database, item *Category) error {
	return db.C("categories").Insert(item)
}

func UpdateCategory(db *mgo.Database, categoryId string, item *Category) error {
	return db.C("categories").UpdateId(bson.ObjectIdHex(categoryId), item)
}
