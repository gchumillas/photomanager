package manager

import "github.com/globalsign/mgo/bson"

// TODO: move this to user_manager.go
// User entity.
type User struct {
	ID          bson.ObjectId   `json:"id" bson:"_id,omitempty"`
	AccessToken string          `json:"accessToken" bson:"accessToken"`
	AccountID   string          `json:"accountId" bson:"accountId"`
	CategoryIDs []bson.ObjectId `json:"categoryIds" bson:"categoryIds"`
}

// TODO: move this to category_manager.go
// Category entity.
type Category struct {
	ID       bson.ObjectId   `json:"id" bson:"_id,omitempty"`
	UserID   bson.ObjectId   `json:"userId" bson:"userId"`
	Name     string          `json:"name"`
	ImageIDs []bson.ObjectId `json:"imageIds" bson:"imageIds"`
}

// QueryOptions options.
type QueryOptions struct {
	Skip  int
	Limit int
	Sort  []string
}
