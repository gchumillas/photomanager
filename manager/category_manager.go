package manager

import (
	"log"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

func (user *User) CreateCategory(db *mgo.Database, cat *Category) {
	cat.ID = bson.NewObjectId()
	cat.UserID = user.ID

	if err := db.C("categories").Insert(cat); err != nil {
		log.Fatal(err)
	}
}

func (user *User) ReadCategory(db *mgo.Database, catID string, cat *Category) (found bool) {
	query := bson.M{"_id": bson.ObjectIdHex(catID), "userId": user.ID}

	if err := db.C("categories").Find(query).One(cat); err != nil {
		switch err {
		case mgo.ErrNotFound:
			return false
		default:
			log.Fatal(err)
		}
	}

	return true
}

func (user *User) UpdateCategory(db *mgo.Database, catID string, cat *Category) (found bool) {
	query := bson.M{"_id": bson.ObjectIdHex(catID), "userId": user.ID}

	if err := db.C("categories").Update(query, cat); err != nil {
		switch err {
		case mgo.ErrNotFound:
			return false
		default:
			log.Fatal(err)
		}
	}

	return true
}

func (user *User) DeleteCategory(db *mgo.Database, catID string) (found bool) {
	query := bson.M{"_id": bson.ObjectIdHex(catID), "userId": user.ID}

	if err := db.C("categories").Remove(query); err != nil {
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
func (user *User) GetCategories(db *mgo.Database, options QueryOptions, parentCatID string, items *[]Category) {
	query := bson.M{"parentCategoryId": bson.ObjectIdHex(parentCatID)}

	if err := db.C("categories").
		Find(query).
		Skip(options.Skip).
		Limit(options.Limit).
		Sort(options.Sort...).
		All(items); err != nil {
		log.Fatal(err)
	}
}

// GetNumCategories returns the number of categories.
func (user *User) GetNumCategories(db *mgo.Database, options QueryOptions, parentCatID string) int {
	query := bson.M{"parentCategoryId": bson.ObjectIdHex(parentCatID)}

	count, err := db.C("categories").Find(query).Count()
	if err != nil {
		log.Fatal(err)
	}

	return count
}
