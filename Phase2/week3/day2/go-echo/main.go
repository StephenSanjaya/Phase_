package main

import (
	"Phase2/week3/day2/go-echo/authorization"
	"Phase2/week3/day2/go-echo/config"
	"Phase2/week3/day2/go-echo/handlers"
	"Phase2/week3/day2/go-echo/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Connect to database
	dsn := "root:@tcp(127.0.0.1:3306)/hacktiv8?charset=utf8mb4&parseTime=True&loc=Local"
	config.InitDatabase(dsn)

	// Migration
	config.DB.AutoMigrate(model.Book{})

	e := echo.New()
	e.Use(middleware.Logger())
	e.POST("/create", handlers.CreateBook)

	b := e.Group("data")
	b.GET("/books", handlers.GetBooks)
	b.GET("/book/:id", handlers.GetBookById)

	c := e.Group("private")
	c.Use(authorization.AuthMiddleware) // implement custom middleware
	c.GET("/balance", func(c echo.Context) error {
		return c.String(http.StatusOK, "Empty")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
