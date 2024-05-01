package main

import (
	pb "Phase3/week2/day2/go-grpc-redis/gateway-service/pb"
	"net/http"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

const (
	userServiceAddress  = "localhost:50051"
	orderServiceAddress = "localhost:50052"
)

func main() {
	e := echo.New()

	// Setup gRPC clients
	userConn, err := grpc.Dial(userServiceAddress, grpc.WithInsecure())
	if err != nil {
		e.Logger.Fatal("Could not connect to User Service: ", err)
	}
	userClient := pb.NewUserServiceClient(userConn)

	orderConn, err := grpc.Dial(orderServiceAddress, grpc.WithInsecure())
	if err != nil {
		e.Logger.Fatal("Could not connect to Order Service: ", err)
	}
	orderClient := pb.NewOrderServiceClient(orderConn)

	// ROUTES API
	e.GET("/users/:id", func(c echo.Context) error {
		userID := c.Param("id")
		res, err := userClient.GetUser(c.Request().Context(), &pb.UserIdRequest{Id: userID})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, res)
	})

	e.POST("/users", func(c echo.Context) error {
		var user pb.UserRequest
		if err := c.Bind(&user); err != nil {
			return err
		}
		res, err := userClient.CreateUser(c.Request().Context(), &user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, res)
	})

	e.GET("/orders/:id", func(c echo.Context) error {
		orderID := c.Param("id")
		res, err := orderClient.GetOrder(c.Request().Context(), &pb.OrderIdRequest{Id: orderID})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, res)
	})

	e.POST("/orders", func(c echo.Context) error {
		var order pb.OrderRequest
		if err := c.Bind(&order); err != nil {
			return err
		}
		res, err := orderClient.CreateOrder(c.Request().Context(), &order)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, res)
	})

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
