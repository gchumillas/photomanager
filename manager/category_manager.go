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

func GetCategories(db *mgo.Database, filter Filter, items *[]Category) {
	if err := db.C("categories").
		Find(filter.Query).
		Skip(filter.Skip).
		Limit(filter.Limit).
		Sort(filter.SortCols...).
		All(items); err != nil {
		log.Fatal(err)
	}
}

func GetNumCategories(db *mgo.Database, filter Filter) int {
	count, err := db.C("categories").Find(filter.Query).Count()
	if err != nil {
		log.Fatal(err)
	}

	return count
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
