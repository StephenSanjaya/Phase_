package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Conn *mongo.Collection
}

var instance *Database
var once sync.Once

func GetConnection() *Database {
	once.Do(func() {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("failed to load env")
			return
		}
		dbName := os.Getenv("DATABASE_NAME")
		dbHost := os.Getenv("DATABASE_HOST")
		dbPort := os.Getenv("DATABASE_PORT")
		dbCollection := os.Getenv("DATABASE_COLLECTION")

		uri := fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
		if err != nil {
			log.Fatal("failed to connect mongo db")
			return
		}

		collection := client.Database(dbName).Collection(dbCollection)
		instance = &Database{Conn: collection}
	})
	return instance
}
