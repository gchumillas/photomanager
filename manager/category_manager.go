package manager

import (
	"log"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

func (cat *Category) CreateCategory(db *mgo.Database, user *User) {
	cat.ID = bson.NewObjectId()
	cat.UserID = user.ID

	if err := db.C("categories").Insert(cat); err != nil {
		log.Fatal(err)
	}
}

func (cat *Category) ReadCategory(db *mgo.Database, user *User, catID string) (found bool) {
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

func (cat *Category) UpdateCategory(db *mgo.Database, user *User, catID string) (found bool) {
	query := bson.M{"_id": bson.ObjectIdHex(catID), "userId": user.ID}
	cat.ID = bson.ObjectIdHex(catID)
	cat.UserID = user.ID

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

// TODO: we do not need cat
func (cat *Category) DeleteCategory(db *mgo.Database, user *User, catID string) (found bool) {
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
