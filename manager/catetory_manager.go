package manager

import (
	"github.com/globalsign/mgo"
)

type Category struct {
	Name string
}

func GetCategories(db *mgo.Database, items *[]Category) error {
	return db.C("categories").Find(nil).All(items)
}

// func GetCategories(db *mgo.Database) (items []Category, err error) {
// 	err = db.C("categories").Find(nil).All(&items)
//
// 	return
// }
