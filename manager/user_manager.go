package manager

import (
	"log"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// User entity.
type User struct {
	ID          bson.ObjectId   `json:"id" bson:"_id,omitempty"`
	AccessToken string          `json:"accessToken" bson:"accessToken"`
	AccountID   string          `json:"accountId" bson:"accountId"`
	CategoryIDs []bson.ObjectId `json:"categoryIds" bson:"categoryIds"`
}

// CreateUser creates a user.
func (user *User) CreateUser(db *mgo.Database) {
	if err := db.C("users").Insert(user); err != nil {
		log.Fatal(err)
	}
}

// ReadUserByAccountID searches a user by Dropbox ID.
func (user *User) ReadUserByAccountID(db *mgo.Database) (found bool) {
	filter := bson.M{"accountId": user.AccountID}

	if err := db.C("users").Find(filter).One(user); err != nil {
		switch err {
		case mgo.ErrNotFound:
			return false
		default:
			log.Fatal(err)
		}
	}

	return true
}

// UpdateUser updates a user.
func (user *User) UpdateUser(db *mgo.Database) (found bool) {
	if err := db.C("users").UpdateId(user.ID, user); err != nil {
		switch err {
		case mgo.ErrNotFound:
			return false
		default:
			log.Fatal(err)
		}
	}

	return true
}
