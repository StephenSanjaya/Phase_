package handlers

import (
	"Phase2/week2/day2/NGC-7/config"
	"Phase2/week2/day2/NGC-7/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllProducts(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")

	rows, err := config.Db.Query("SELECT product_id, name, description, image_url, price, store_id FROM Products")
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorMessage{
			Status:  http.StatusInternalServerError,
			Message: "Failed to get all products" + err.Error(),
		})
		return
	}

	var products []models.Product

	for rows.Next() {
		var p models.Product
		err = rows.Scan(&p.ProductID, &p.Name, &p.Description, &p.ImageUrl, &p.Price, &p.StoreID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorMessage{
				Status:  http.StatusInternalServerError,
				Message: "Failed to get all products" + err.Error(),
			})
			return
		}
		products = append(products, p)
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

	statement, err := config.Db.Prepare("SELECT product_id, name, description, image_url, price, store_id FROM Products WHERE product_id = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorMessage{
			Status:  http.StatusInternalServerError,
			Message: "Failed to preare query, " + err.Error(),
		})
		return
	}
	defer statement.Close()

	var p models.Product

	err = statement.QueryRow(id).Scan(&p.ProductID, &p.Name, &p.Description, &p.ImageUrl, &p.Price, &p.StoreID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorMessage{
			Status:  http.StatusInternalServerError,
			Message: "Failed to scan query, " + err.Error(),
		})
		return
	}
	defer statement.Close()

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

	statement, err := config.Db.Prepare("INSERT INTO Products (name, description, image_url,  price, store_id) VALUES (?,?,?,?,?)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorMessage{
			Status:  http.StatusInternalServerError,
			Message: "Failed to prepare query, " + err.Error(),
		})
		return
	}
	defer statement.Close()

	_, err = statement.Exec(p.Name, p.Description, p.ImageUrl, p.Price, p.StoreID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorMessage{
			Status:  http.StatusInternalServerError,
			Message: "Failed to insert new data, " + err.Error(),
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

	statement, err := config.Db.Prepare("UPDATE Products SET name = ?, description = ?, image_url = ?, price = ?, store_id = ? WHERE product_id = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorMessage{
			Status:  http.StatusInternalServerError,
			Message: "Failed to prepare query, " + err.Error(),
		})
		return
	}
	defer statement.Close()

	_, err = statement.Exec(p.Name, p.Description, p.ImageUrl, p.Price, p.StoreID, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorMessage{
			Status:  http.StatusInternalServerError,
			Message: "Failed to insert id not found, " + err.Error(),
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

	statement, err := config.Db.Prepare("DELETE FROM Products WHERE product_id = ?")
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorMessage{
			Status:  http.StatusInternalServerError,
			Message: "Failed to prepare query, " + err.Error(),
		})
		return
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorMessage{
			Status:  http.StatusInternalServerError,
			Message: "Failed to delete id not found, " + err.Error(),
		})
		return
	}

	var success_msg = models.SuccessMessage{
		Status:  http.StatusOK,
		Message: "Success delete data by id",
	}

	c.JSON(http.StatusOK, success_msg)
}
