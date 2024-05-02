package main

import (
	"Phase3/week2/day2/NGC-5/ms-gateaway/controller"
	pb "Phase3/week2/day2/NGC-5/ms-gateaway/pb/auth"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
)

const (
	userServiceAddress = "localhost:50051"
)

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	userConn, err := grpc.Dial(userServiceAddress, grpc.WithInsecure())
	if err != nil {
		e.Logger.Fatal("Could not connect to User Service: ", err)
	}
	userClient := pb.NewAuthServiceClient(userConn)

	c := controller.NewAuthController(userClient)

	e.POST("/users/register", c.RegisterAuth)
	e.GET("/users", c.GetUsers)

	e.Logger.Fatal(e.Start(":8081"))
}
