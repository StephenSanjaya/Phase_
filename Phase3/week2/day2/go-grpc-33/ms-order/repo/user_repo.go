package repo

import (
	"Phase3/week2/day2/go-grpc-33/ms-order/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	Collection *mongo.Collection
}

func NewUserRepo(collection *mongo.Collection) UserRepo {
	return UserRepo{
		Collection: collection,
	}
}

func (ur *UserRepo) Insert(data models.User) error {
	_, err := ur.Collection.InsertOne(context.TODO(), data)
	if err != nil {
		fmt.Println("ERR 2", err)
		return err
	}

	return nil
}
