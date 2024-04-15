package main

import (
	"Phase2/week4/day1/preview-week3/config"
	"Phase2/week4/day1/preview-week3/handler"
	"Phase2/week4/day1/preview-week3/middleware"

	"github.com/labstack/echo/v4"
)

func main() {
	db := config.GetConnection()

	authService := handler.NewAuthService(db)
	playerService := handler.NewPlayerService(db)

	e := echo.New()
	e.Use(middleware.MiddlewareLogging)
	e.HTTPErrorHandler = middleware.ErrorHandler

	auth := e.Group("/auth")
	{
		auth.POST("/login", authService.LoginHandler)
		auth.POST("/refresh", authService.RefreshTokenHandler)
	}

	players := e.Group("/players")
	players.Use(middleware.AuthMiddleware)
	{
		players.GET("", playerService.GetAllPlayers)
		players.POST("", playerService.CreateNewPlayers)
		players.PUT("/:id", playerService.UpdateNewPlayers)
	}

	e.Logger.Fatal(e.Start(":8081"))
}
