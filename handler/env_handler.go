package handler

import (
	"github.com/globalsign/mgo"
)

type Env struct {
	session *mgo.Session
}

func NewEnv(session *mgo.Session) *Env {
	return &Env{session}
}

// Gets the selected Database.
func (env *Env) DB() *mgo.Database {
	return env.session.DB("")
}
