package models

import (
	"gopkg.in/mgo.v2/bson"
)

// User model
type User struct {
	Name   string        `json:"name" bson:"name"`
	Gender string        `json:"gender" bson:"gender"`
	Age    int           `json:"age" bson:"age"`
	ID     bson.ObjectId `json:"id" bson:"_id"`
}
