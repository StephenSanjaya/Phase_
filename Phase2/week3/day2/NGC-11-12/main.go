package main

import (
	"Phase2/week3/day2/NGC-11-12/config"
	"Phase2/week3/day2/NGC-11-12/handlers"
	"Phase2/week3/day2/NGC-11-12/middleware"

	_ "Phase2/week3/day2/NGC-11-12/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title           NGC11-12 Api Docs
// @version         1.0
// @description     This is ngc11-12 api docs

// @contact.name   stephen
// @contact.email  stephen@email.com

// @host      localhost:8081
// @BasePath  /
func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/hc8?charset=utf8mb4&parseTime=True&loc=Local"
	config.InitDatabase(dsn)

	e := echo.New()
	e.Use(middleware.MiddlewareLogging)
	e.HTTPErrorHandler = middleware.ErrorHandler

	users := e.Group("/users")
	users.POST("/register", handlers.RegisterHandler)
	users.POST("/login", handlers.LoginHandler)

	e.GET("/products", middleware.AuthMiddleware(handlers.GetAllProducts))

	e.POST("/transactions", middleware.AuthMiddleware(handlers.CreateNewTransaction))

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8081"))
}
