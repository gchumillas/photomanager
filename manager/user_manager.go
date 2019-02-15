package manager

import (
	"log"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// CreateUser creates a user.
func (user *User) CreateUser(db *mgo.Database) {
	if err := db.C("users").Insert(user); err != nil {
		log.Fatal(err)
	}
}

// ReadUserByAccountID searches a user by account ID.
func (user *User) ReadUserByAccountID(db *mgo.Database) (found bool) {
	query := bson.M{"accountId": user.AccountID}

	if err := db.C("users").Find(query).One(user); err != nil {
		switch err {
		case mgo.ErrNotFound:
			return false
		default:
			log.Fatal(err)
		}
	}

	return true
}

// ReadUserByAccountID searches a user by access token.
func (user *User) ReadUserByToken(db *mgo.Database) (found bool) {
	query := bson.M{"accessToken": user.AccessToken}

	if err := db.C("users").Find(query).One(user); err != nil {
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
