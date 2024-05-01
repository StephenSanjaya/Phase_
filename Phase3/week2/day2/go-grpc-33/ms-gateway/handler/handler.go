package handler

import (
	"Phase3/week2/day2/go-grpc-33/ms-gateway/model"
	pb "Phase3/week2/day2/go-grpc-33/ms-gateway/pb"
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	userGRPC pb.UserServiceClient
}

func NewHandler(userGRPC pb.UserServiceClient) *Handler {
	return &Handler{
		userGRPC: userGRPC,
	}
}

func (h *Handler) CreateUser(c echo.Context) error {
	var payload model.User

	err := c.Bind(&payload)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	in := pb.User{
		Age:     payload.Age,
		Address: payload.Address,
		Name:    payload.Name,
	}

	response, err := h.userGRPC.CreateUser(context.TODO(), &in)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, response)
}
