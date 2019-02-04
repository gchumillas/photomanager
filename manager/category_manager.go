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

// CreateCategory inserts a category.
func (cat *Category) CreateCategory(db *mgo.Database) {
	if err := db.C("categories").Insert(cat); err != nil {
		log.Fatal(err)
	}
}

// ReadCategory return a category.
func (cat *Category) ReadCategory(db *mgo.Database) (found bool) {
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

// UpdateCategory updates a category.
// TODO: found bool --> bool
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

// DeleteCategory deletes a category.
func (cat *Category) DeleteCategory(db *mgo.Database) bool {
	if err := db.C("categories").RemoveId(cat.ID); err != nil {
		switch err {
		case mgo.ErrNotFound:
			return false
		default:
			log.Fatal(err)
		}
	}

	return true
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
