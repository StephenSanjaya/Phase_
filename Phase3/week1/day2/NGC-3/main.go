package main

import (
	"Phase3/week1/day2/NGC-3/src/config"
	"Phase3/week1/day2/NGC-3/src/controller"
	"Phase3/week1/day2/NGC-3/src/middleware"
	"Phase3/week1/day2/NGC-3/src/repository"
	"Phase3/week1/day2/NGC-3/src/router"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	coll := config.GetConnection().Conn

	validate := validator.New()
	msgRepo := repository.NewMessageRepository(coll)
	msgController := controller.NewMessageController(msgRepo, validate)

	e := echo.New()
	e.Use(middleware.MiddlewareLogging)
	e.HTTPErrorHandler = middleware.ErrorHandler

	router.SetupRouter(e, msgController)

	e.Logger.Fatal(e.Start(":8081"))
}
