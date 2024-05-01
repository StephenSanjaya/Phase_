package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id      primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	Name    string             `bson:"name, omitempty" json:"name"`
	Age     string             `bson:"age, omitempty" json:"age"`
	Address string             `bson:"address, omitempty" json:"address"`
}
