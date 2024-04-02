package main

import (
	"Phase2/week3/day1/latihan-slc/config"
	"Phase2/week3/day1/latihan-slc/handler"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := config.InitDb()
	r := gin.Default()

	dbHandler := handler.NewDbHandler(db)
	customerController := handler.NewCustomerHandler(dbHandler)
	productController := handler.NewProductHandler(dbHandler)

	r.POST("/register", customerController.Register)
	r.POST("/login", customerController.Login)

	r.GET("/products", productController.GetAllProducts)

	r.Run(":8080")
}
