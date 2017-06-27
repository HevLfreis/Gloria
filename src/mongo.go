package main

import (
	"gopkg.in/mgo.v2"
)

type MongoStore struct {
	session *mgo.Session
}

func NewMongoStore() *MongoStore {
	session, err := mgo.Dial(MONGO_URL)
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

//func main()  {
//	session, err := mgo.Dial("mongodb://localhost")
//	if err != nil {
//		panic(err)
//	}
//	defer session.Close()
//
//	// Optional. Switch the session to a monotonic behavior.
//	c := session.DB("chat2x").C("messages")
//
//}
