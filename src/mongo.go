package main

import (
	"gopkg.in/mgo.v2"
)

type MongoStore struct {
	session *mgo.Session
}

func NewMongoStore(url string) *MongoStore {
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return &MongoStore{session}
}

func (ms *MongoStore) Session() *mgo.Session {
	return ms.session.Clone()
}

func (ms *MongoStore) Close() {
	ms.session.Close()
}
