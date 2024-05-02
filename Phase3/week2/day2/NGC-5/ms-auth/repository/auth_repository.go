package repository

import (
	"Phase3/week2/day2/NGC-5/ms-auth/model"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type AuthRepository struct {
	db *mongo.Collection
}

func NewAuthRepository(db *mongo.Collection) *AuthRepository {
	return &AuthRepository{db: db}
}

func (ar *AuthRepository) Insert(user *model.User) (*model.User, error) {
	res, err := ar.db.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	user.Id = res.InsertedID.(primitive.ObjectID)

	return user, nil
}

func (ar *AuthRepository) FindAll() (*[]model.User, error) {
	users := new([]model.User)

	cursor, err := ar.db.Find(context.TODO(), bson.M{})
	if err != nil {
		return users, err
	}

	if err := cursor.All(context.TODO(), users); err != nil {
		return users, err
	}

	return users, nil
}
