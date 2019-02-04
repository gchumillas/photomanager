package manager

import (
	"log"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// Category entity.
type Category struct {
	ID       bson.ObjectId   `json:"id" bson:"_id,omitempty"`
	Name     string          `json:"name"`
	ImageIDs []bson.ObjectId `json:"imageIds" bson:"imageIds"`
}

// GetCategories returns a list of categories..
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

// GetNumCategories returns the number of categories.
func GetNumCategories(db *mgo.Database, filter Filter) int {
	count, err := db.C("categories").Find(filter.Query).Count()
	if err != nil {
		log.Fatal(err)
	}

	return count
}

// GetCategory return a category.
// TODO: rename by ReadCategory
func (cat *Category) GetCategory(db *mgo.Database) (found bool) {
	if err := db.C("categories").FindId(cat.ID).One(cat); err != nil {
		switch err {
		case mgo.ErrNotFound:
			return false
		default:
			log.Fatal(err)
		}
	}

	return true
}

// InsertCategory inserts a category.
// TODO: rename by CreateCategory
// TODO: must return the ID
func (cat *Category) InsertCategory(db *mgo.Database) {
	if err := db.C("categories").Insert(cat); err != nil {
		log.Fatal(err)
	}
}

// UpdateCategory updates a category.
func (cat *Category) UpdateCategory(db *mgo.Database) (found bool) {
	if err := db.C("categories").UpdateId(cat.ID, cat); err != nil {
		switch err {
		case mgo.ErrNotFound:
			return false
		default:
			log.Fatal(err)
		}
	}

	return true
}
