package main

import (
	"Phase2/week3/day2/NGC-11-12-13/config"
	"Phase2/week3/day2/NGC-11-12-13/handlers"
	"Phase2/week3/day2/NGC-11-12-13/middleware"

	_ "Phase2/week3/day2/NGC-11-12-13/docs"

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
	// goose -dir migrations mysql root:@/hc8?multiStatements=true"&"parseTime=true up
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

	stores := e.Group("/stores")
	stores.Use(middleware.AuthMiddleware)
	{
		stores.GET("", handlers.GetAllStores)
		stores.GET("/:id", handlers.GetStoreById)
	}

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8081"))
}
