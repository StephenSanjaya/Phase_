package repository

import (
	"Phase3/week1/day2/NGC-3/src/model"
	"Phase3/week1/day2/NGC-3/src/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageRepositoryI interface {
	Save(message *model.Message) *utils.HTTPError
	FindByID(id primitive.ObjectID) (*model.Message, *utils.HTTPError)
	FindAllMessageBySender(sender_name string) (*[]model.Message, *utils.HTTPError)
}
