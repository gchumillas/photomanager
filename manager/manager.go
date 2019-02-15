package manager

import "github.com/globalsign/mgo/bson"

// User entity.
type User struct {
	ID          bson.ObjectId   `json:"id" bson:"_id,omitempty"`
	AccessToken string          `json:"accessToken" bson:"accessToken"`
	AccountID   string          `json:"accountId" bson:"accountId"`
	CategoryIDs []bson.ObjectId `json:"categoryIds" bson:"categoryIds"`
}

// Category entity.
type Category struct {
	ID       bson.ObjectId   `json:"id" bson:"_id,omitempty"`
	Name     string          `json:"name"`
	ImageIDs []bson.ObjectId `json:"imageIds" bson:"imageIds"`
}

// QueryOptions options.
type QueryOptions struct {
	Skip     int
	Limit    int
	Query    interface{}
	SortCols []string
}
