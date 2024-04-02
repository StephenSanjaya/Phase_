package main

import (
	"Phase2/week2/day5/preview-week1/config"
	"Phase2/week2/day5/preview-week1/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.GetConnection()

	authService := handlers.NewAuthService(db)

	r := gin.Default()

	api := r.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.POST("/register", authService.RegisterHandler)
		}
	}

	r.Run(":8081")
}
