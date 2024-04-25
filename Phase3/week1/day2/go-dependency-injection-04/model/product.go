package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID   primitive.ObjectID `json:"id"`
	Name string             `json:"name"`
}
