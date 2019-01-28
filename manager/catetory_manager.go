package manager

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type Category struct {
	ID   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name"`
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
