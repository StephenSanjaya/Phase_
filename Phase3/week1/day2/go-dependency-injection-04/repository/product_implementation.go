package repository

import (
	"Phase3/week1/day2/go-dependency-injection-04/config/utils"
	"Phase3/week1/day2/go-dependency-injection-04/model"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductImplementation struct {
	DB *mongo.Collection
}

func NewProductRepository(db *mongo.Collection) ProductImplementation {
	return ProductImplementation{
		DB: db,
	}
}

func (p ProductImplementation) Create(data *model.Product) *utils.ErrResponse {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := p.DB.InsertOne(ctx, data)
	if err != nil {
		errResponse := &utils.ErrInternalServer
		errResponse.Detail = err.Error()
		return errResponse
	}
	data.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}
