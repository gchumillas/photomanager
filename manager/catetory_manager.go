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
