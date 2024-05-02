package config

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	User *mongo.Collection
}

var instance *Database
var once sync.Once

func GetConnection() *Database {
	once.Do(func() {
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
			log.Fatal("failed to connect mongo db")
			return
		}

		userColl := client.Database("ngc5-7").Collection("users")
		instance = &Database{User: userColl}
	})

	return instance
}
