package manager

import (
	"log"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

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
func (cat *Category) DeleteCategory(db *mgo.Database) (found bool) {
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
func (user *User) GetCategories(db *mgo.Database, options QueryOptions, items *[]Category) {
	query := bson.M{
		"parentCategoryId": nil,
	}

	if err := db.C("categories").
		Find(query).
		Skip(options.Skip).
		Limit(options.Limit).
		Sort(options.SortCols...).
		All(items); err != nil {
		log.Fatal(err)
	}
}

// GetSubategories returns a list of categories..
func (user *User) GetSubcategories(db *mgo.Database, options QueryOptions, parentCatID string, items *[]Category) {
	query := bson.M{
		"parentCategoryId": bson.ObjectIdHex(parentCatID),
	}

	if err := db.C("categories").
		Find(query).
		Skip(options.Skip).
		Limit(options.Limit).
		Sort(options.SortCols...).
		All(items); err != nil {
		log.Fatal(err)
	}
}

// GetNumCategories returns the number of categories.
func (user *User) GetNumCategories(db *mgo.Database, options QueryOptions) int {
	query := bson.M{
		"parentCategoryId": nil,
	}

	count, err := db.C("categories").Find(query).Count()
	if err != nil {
		log.Fatal(err)
	}

	return count
}

// GetNumCategories returns the number of categories.
func (user *User) GetNumSubcategories(db *mgo.Database, options QueryOptions, parentCatID string) int {
	query := bson.M{
		"parentCategoryId": bson.ObjectIdHex(parentCatID),
	}

	count, err := db.C("categories").Find(query).Count()
	if err != nil {
		log.Fatal(err)
	}

	return count
}
