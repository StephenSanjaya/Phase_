package repository

import (
	"Phase3/week1/day2/NGC-3/src/model"
	"Phase3/week1/day2/NGC-3/src/utils"
	"context"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageRepositoryImpl struct {
	db *mongo.Collection
}

func NewMessageRepository(db *mongo.Collection) MessageRepositoryI {
	return &MessageRepositoryImpl{db: db}
}

func (mi *MessageRepositoryImpl) Save(message *model.Message) *utils.HTTPError {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := mi.db.InsertOne(ctx, message)
	if err != nil {
		return &utils.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "failed to insert",
			Detail:  err,
		}
	}
	message.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

func (mi *MessageRepositoryImpl) FindByID(message_id primitive.ObjectID) (*model.Message, *utils.HTTPError) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	msg := new(model.Message)

	err := mi.db.FindOne(ctx, bson.M{"_id": message_id}).Decode(&msg)
	if err != nil {
		return msg, &utils.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "failed to insert",
			Detail:  err,
		}
	}

	return msg, nil
}

func (mi *MessageRepositoryImpl) FindAllMessageBySender(sender_name string) (*[]model.Message, *utils.HTTPError) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	msgs := new([]model.Message)

	cursor, err := mi.db.Find(ctx, bson.M{"user_sender.name": sender_name})
	if err != nil {
		return msgs, &utils.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "failed to insert",
			Detail:  err,
		}
	}
	for cursor.Next(context.Background()) {
		var msg model.Message
		if err := cursor.Decode(&msg); err != nil {
			return msgs, &utils.HTTPError{
				Code:    http.StatusInternalServerError,
				Message: "failed to decode",
				Detail:  err,
			}
		}
		*msgs = append(*msgs, msg)
	}

	if err := cursor.Err(); err != nil {
		return msgs, &utils.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: "mongo cursor failed",
			Detail:  err,
		}
	}

	return msgs, nil
}
