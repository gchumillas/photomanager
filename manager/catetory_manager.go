package manager

import (
	"github.com/globalsign/mgo"
)

type Category struct {
	name string
}

func GetCategories(db *mgo.Database) ([]Category, error) {
	// categories := db.Collection("categories")
	//
	// cur, err := categories.Find(context.Background(), nil)
	// defer cur.Close(context.Background())
	//
	// for cur.Next(context.Background()) {
	// 	// do something
	// }
	//
	// return nil, err

	return nil, nil
}
