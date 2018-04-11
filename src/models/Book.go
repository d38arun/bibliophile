package models

import "gopkg.in/mgo.v2/bson"

type Book struct {
	ID bson.ObjectId `bson:"_id" json:"id"`
	Title string `bson:"title" json:"title"`
	Author string `bson:"author" json:"author"`
	Description string `bson:"description" json:"description"`
	CoverImage string `bson:"cover_image" json:"cover_image"`
}