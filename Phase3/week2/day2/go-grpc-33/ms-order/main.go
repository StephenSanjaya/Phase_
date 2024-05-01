package main

import (
	"Phase3/week2/day2/go-grpc-33/ms-order/cmd"
	"Phase3/week2/day2/go-grpc-33/ms-order/config"
	"Phase3/week2/day2/go-grpc-33/ms-order/handler"
	"Phase3/week2/day2/go-grpc-33/ms-order/repo"
)

func main() {
	conn := config.MongoConnect().Database("hacktiv8")
	userCollection := conn.Collection("users")

	userRepo := repo.NewUserRepo(userCollection)
	userHandler := handler.NewHandler(userRepo)
	cmd.InitGrpc(userHandler)
}
