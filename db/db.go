package db

import (
	mgo "gopkg.in/mgo.v2"
)

type Db struct {
	Session    *mgo.Session
	Collection *mgo.Collection
}

func Init() Db {
	// Connect to the MongoDB database.
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	// defer session.Close()

	collection := GetCollection("fitgirldb", "games", session)
	mydb := Db{Session: session, Collection: collection}
	return mydb
}

func GetCollection(dbname, collectionname string, session *mgo.Session) *mgo.Collection {
	// Get the Collection.
	return session.DB(dbname).C(collectionname)
}
