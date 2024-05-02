package controller

import (
	pb "Phase3/week2/day2/NGC-5/ms-gateaway/pb/auth"
	"net/http"

	"github.com/labstack/echo/v4"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type AuthController struct {
	authGRPC pb.AuthServiceClient
}

func NewAuthController(authGRPC pb.AuthServiceClient) *AuthController {
	return &AuthController{authGRPC: authGRPC}
}

func (ac *AuthController) RegisterAuth(c echo.Context) error {
	user := new(pb.UserRequest)

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := ac.authGRPC.RegisterAuth(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, res)
}

func (ac *AuthController) GetUsers(c echo.Context) error {
	res, err := ac.authGRPC.GetUsers(c.Request().Context(), &emptypb.Empty{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}
