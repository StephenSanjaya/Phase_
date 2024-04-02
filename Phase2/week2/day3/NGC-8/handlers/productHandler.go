package handlers

import (
	"Phase2/week2/day3/NGC-8/config"
	"Phase2/week2/day3/NGC-8/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllProducts(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")

	var products []models.Product
	if res := config.Db.Find(&products); res.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorMessage{
			Status:  http.StatusInternalServerError,
			Message: "Failed to get all products: " + res.Error.Error(),
		})
		return
	}

	var success_message = models.SuccessMessage{
		Status:  http.StatusOK,
		Message: "Success to get all products",
		Datas:   products,
	}

	c.JSON(http.StatusOK, success_message)
}

func GetProductById(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	id := c.Param("id")

	var p models.Product
	if res := config.Db.First(&p, id); res.Error != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorMessage{
			Status:  http.StatusInternalServerError,
			Message: "Id not found: " + res.Error.Error(),
		})
		return
	}

	var success_msg = models.SuccessMessage{
		Status:  http.StatusOK,
		Message: "Success get data by id",
		Datas:   p,
	}

	c.JSON(http.StatusOK, success_msg)
}

func InsertNewProduct(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")

	var p models.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorMessage{
			Status:  http.StatusBadRequest,
			Message: "invalid body input: " + err.Error(),
		})
		return
	}

	if res := config.Db.Create(&p); res.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorMessage{
			Status:  http.StatusInternalServerError,
			Message: "Failed to create new product: " + res.Error.Error(),
		})
		return
	}

	var success_msg = models.SuccessMessage{
		Status:  http.StatusCreated,
		Message: "Success insert new data",
	}

	c.JSON(http.StatusCreated, success_msg)
}

func Updateproduct(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")

	id := c.Param("id")

	var p models.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorMessage{
			Status:  http.StatusBadRequest,
			Message: "invalid body input: " + err.Error(),
		})
		return
	}

	if res := config.Db.Model(models.Product{}).Where("product_id = ?", id).Updates(&p); res.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorMessage{
			Status:  http.StatusInternalServerError,
			Message: "Failed to create new product: " + res.Error.Error(),
		})
		return
	}

	var success_msg = models.SuccessMessage{
		Status:  http.StatusOK,
		Message: "Success update data by id",
	}

	c.JSON(http.StatusBadRequest, success_msg)
}

func DeleteProduct(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")

	id := c.Param("id")

	var product models.Product
	if res := config.Db.Delete(&product, id); res.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorMessage{
			Status:  http.StatusInternalServerError,
			Message: "Failed to create new product: " + res.Error.Error(),
		})
		return
	}

	var success_msg = models.SuccessMessage{
		Status:  http.StatusOK,
		Message: "Success delete data by id",
	}

	c.JSON(http.StatusOK, success_msg)
}
