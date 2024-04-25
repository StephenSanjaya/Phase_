package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Message      string             `bson:"message" json:"message" validate:"required"`
	UserReceiver UserReceiver       `bson:"user_receiver" json:"user_receiver"`
	UserSender   UserSender         `bson:"user_sender" json:"user_sender"`
}

type UserReceiver struct {
	PhoneNumber string `bson:"phone_number" json:"phone_number"`
	Name        string `bson:"name" json:"name"`
}
type UserSender struct {
	PhoneNumber string `bson:"phone_number" json:"phone_number" validate:"required"`
	Name        string `bson:"name" json:"name" validate:"required"`
}
