package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	DbHandler
}

func NewProductHandler(dbHandler DbHandler) ProductHandler{
	return ProductHandler{
		DbHandler: dbHandler,
	}
}

func (ph ProductHandler) GetAllProducts(c *gin.Context) {
	products, dbErr := ph.DbHandler.FindAllProductsInDb()
	if dbErr != nil {
		ErrJsonWriter(c, *dbErr, "Failed to get products from database")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Get all products",
		"data": products,
	})
}