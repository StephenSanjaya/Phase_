package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

// Person struct
type Employee struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Name    string             `json:"name"`
	Age     int                `json:"age"`
	Address string             `json:"address"`
	Salary  float64            `json:"salary"`
}
