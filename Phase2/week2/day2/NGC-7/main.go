package main

import (
	"Phase2/week2/day2/NGC-7/config"
	"Phase2/week2/day2/NGC-7/handlers"
	"Phase2/week2/day2/NGC-7/middleware"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("err loading env")
		return
	}

	var dbConfig config.DBEnv
	err = envconfig.Process("DATABASE", &dbConfig)
	if err != nil {
		log.Fatal("err process env")
		return
	}

	source := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", dbConfig.DBUsername, dbConfig.DBPassword, dbConfig.DBHost, dbConfig.DBPort, dbConfig.DBName)

	config.GetConnection(source)
	defer config.Db.Close()

	r := gin.Default()
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
