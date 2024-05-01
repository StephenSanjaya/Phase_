package main

import (
	"Phase3/week2/day2/go-grpc-33/ms-gateway/handler"
	pb "Phase3/week2/day2/go-grpc-33/ms-gateway/pb"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello world")
	grpcConn, err := grpc.Dial("50001", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	grpcService := pb.NewUserServiceClient(grpcConn)
	h := handler.NewHandler(grpcService)

	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// api
	e.POST("/create", h.CreateUser)
	e.Logger.Fatal(e.Start(":1323"))
}
