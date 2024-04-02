package router

import (
	"Phase2/week2/day4/go-swagger/controller"
	"Phase2/week2/day4/go-swagger/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Menggunakan grup berdasarkan versi API (v1)
	v1 := r.Group("/v1")
	{
		// Item API Routes
		itemGroup := v1.Group("/items")
		{
			itemGroup.GET("/", controller.GetAllItems)
			itemGroup.GET("/:id", controller.GetItemByID)
			itemGroup.POST("/", controller.CreateItem)
			itemGroup.PUT("/:id", controller.UpdateItem)
			itemGroup.DELETE("/:id", controller.DeleteItem)
		}
	}
	docs.SwaggerInfo.BasePath = "/v1"
	r.GET("/swagger/*any", ginSwagger.
		WrapHandler(swaggerFiles.Handler))
	return r
}
