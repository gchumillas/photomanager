package manager

import "github.com/globalsign/mgo/bson"

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
