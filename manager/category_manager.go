package manager

import (
	"log"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

func NewCategory(catID ...string) *Category {
	var id bson.ObjectId
	if len(catID) > 0 {
		id = bson.ObjectIdHex(catID[0])
	}

	return &Category{ID: id}
}

func (cat *Category) CreateCategory(db *mgo.Database, user *User) {
	cat.ID = bson.NewObjectId()
	cat.UserID = user.ID

	if err := db.C("categories").Insert(cat); err != nil {
		log.Fatal(err)
	}
}

func (cat *Category) ReadCategory(db *mgo.Database, user *User) (found bool) {
	query := bson.M{"_id": cat.ID, "userId": user.ID}

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

func (cat *Category) UpdateCategory(db *mgo.Database, user *User) (found bool) {
	query := bson.M{"_id": cat.ID, "userId": user.ID}
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

func (cat *Category) DeleteCategory(db *mgo.Database, user *User) (found bool) {
	query := bson.M{"_id": cat.ID, "userId": user.ID}

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
