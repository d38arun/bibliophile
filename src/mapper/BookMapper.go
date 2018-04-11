package mapper

import (
	"gopkg.in/mgo.v2"
	"log"
	. "../models"
	"gopkg.in/mgo.v2/bson"
)

type BookDAO struct {
	Server string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "Books"
)

func (b *BookDAO) Connect() {
	session, err := mgo.Dial(b.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(b.Database)
}

func (b *BookDAO) FindAllBooks() ([]Book, error) {
	var books []Book
	err := db.C(COLLECTION).Find(bson.M{}).All(&books)
	return books, err
}