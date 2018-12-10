package mongo

import (
	mgo "gopkg.in/mgo.v2"
)

type Database struct {
	Name 	string
	Session *mgo.Session
}