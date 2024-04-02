package main

import (
	"Phase2/week3/day2/NGC-11-12/config"
	"Phase2/week3/day2/NGC-11-12/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/hc8?charset=utf8mb4&parseTime=True&loc=Local"
	config.InitDatabase(dsn)

	e := echo.New()
	e.Use(middleware.Logger())
	users := e.Group("/users")
	users.POST("/register", handlers.RegisterHandler)
	users.POST("/login", handlers.LoginHandler)

	e.GET("/products", handlers.GetAllProducts)

	e.POST("/transactions", handlers.CreateNewTransaction)

	e.Logger.Fatal(e.Start(":8081"))
}
