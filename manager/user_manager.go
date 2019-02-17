package manager

import (
	"log"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// User entity. The AccessToken and AccountID properties are provided by the
// user's Dropbox account.
type User struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	AccessToken string        `json:"accessToken" bson:"accessToken"`
	AccountID   string        `json:"accountId" bson:"accountId"`
}

// NewUser creates a user.
func NewUser(userID ...string) *User {
	var id bson.ObjectId
	if len(userID) > 0 {
		id = bson.ObjectIdHex(userID[0])
	}

	return &User{ID: id}
}

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

// ReadUserByToken searches a user by access token.
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

// GetCategories returns a list of categories..
func (user *User) GetCategories(db *mgo.Database, options QueryOptions, parentCatID string) []Category {
	items := []Category{}

	query := bson.M{"parentCategoryId": nil}
	if len(parentCatID) > 0 {
		query["parentCategoryId"] = bson.ObjectIdHex(parentCatID)
	}

	if err := db.C("categories").
		Find(query).
		Skip(options.Skip).
		Limit(options.Limit).
		Sort(options.Sort...).
		All(&items); err != nil {
		log.Fatal(err)
	}

	return items
}

// GetNumCategories returns the number of categories.
func (user *User) GetNumCategories(db *mgo.Database, options QueryOptions, parentCatID string) int {
	query := bson.M{"parentCategoryId": nil}
	if len(parentCatID) > 0 {
		query["parentCategoryId"] = bson.ObjectIdHex(parentCatID)
	}

	count, err := db.C("categories").Find(query).Count()
	if err != nil {
		log.Fatal(err)
	}

	return count
}
