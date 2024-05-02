package main

import (
	"Phase3/week2/day3/go-pooling-message-broker/internal/app"
	"Phase3/week2/day3/go-pooling-message-broker/internal/db"
	"Phase3/week2/day3/go-pooling-message-broker/internal/mq"
	"log"
)

func main() {
	database, err := db.NewDatabase("postgres://postgres:admin@localhost:5432/ecommerce?pool_max_conns=10") // max 10 orang dalam 1 waktu
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.Close()

	messageQueue, err := mq.NewMessageQueue("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ: %v", err)
	}
	defer messageQueue.Close()

	application := app.NewApp(database, messageQueue)
	application.Start(":8081")
}
