package controller

import "github.com/labstack/echo/v4"

type MessageControllerI interface {
	CreateMessage(c echo.Context) error
	FindMessageByID(c echo.Context) error
	FindAllMessageBySender(c echo.Context) error
}
