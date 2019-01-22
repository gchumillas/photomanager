package manager

import (
	"context"

	"github.com/mongodb/mongo-go-driver/mongo"
)

type Category struct {
	name string
}

func GetAllCategories(db *mongo.Database) ([]Category, error) {
	categories := db.Collection("categories")

	cur, err := categories.Find(context.Background(), nil)
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		// do something
	}

	return nil, err
}
