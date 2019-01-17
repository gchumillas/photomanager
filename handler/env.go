package handler

import "github.com/mongodb/mongo-go-driver/mongo"

type Env struct {
	client *mongo.Client
}

func NewEnv(client *mongo.Client) *Env {
	return &Env{client}
}
