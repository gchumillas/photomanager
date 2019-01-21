package handler

import "github.com/mongodb/mongo-go-driver/mongo"

type Env struct {
	db *mongo.Database
}

func NewEnv(db *mongo.Database) *Env {
	return &Env{db}
}
