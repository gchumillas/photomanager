package handler

import (
	"github.com/globalsign/mgo"
)

type Env struct {
	db *mgo.Database
}

func NewEnv(db *mgo.Database) *Env {
	return &Env{db}
}
