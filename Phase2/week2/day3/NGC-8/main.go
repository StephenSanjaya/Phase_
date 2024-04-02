package main

import (
	"Phase2/week2/day3/NGC-8/config"
	"Phase2/week2/day3/NGC-8/handlers"
	"Phase2/week2/day3/NGC-8/middleware"

	"github.com/gin-gonic/gin"
)

func main() {

	config.GetConnection()

	r := gin.Default()
	r.Use(middleware.ErrorMiddleware)

	users := r.Group("/users")
	{
		users.POST("/login", handlers.LoginHandler)
		users.POST("/register", handlers.RegisterHandler)
	}

	products := r.Group("/products")
	products.Use(middleware.AuthMiddleware())
	{
		products.GET("/", handlers.GetAllProducts)
		products.GET("/:id", handlers.GetProductById)
		products.POST("/", handlers.InsertNewProduct)
		products.PUT("/:id", handlers.Updateproduct)
		products.DELETE("/:id", handlers.DeleteProduct)
	}

	r.Run(":8081")

}
