package controller

import (
	"Phase2/week2/day4/go-swagger/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Item godoc
// @Summary Create items
// @Description Create new items
// @ID Create-items
// @Accept  json
// @Produce  json
// @Param item body utils.Item true "Item to create"
// @Success 201 {object} utils.Item
// @Router /items/ [post]
func CreateItem(c *gin.Context) {
	var item utils.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := utils.CreateItem(item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// Item godoc
// @Summary Get an item by ID
// @Description Get an item by ID
// @ID Get-item-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Item ID"
// @Success 200 {object} utils.Item
// @Router /items/{id} [get]
func GetItemByID(c *gin.Context) {
	id := c.Param("id")
	itemID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	item, err := utils.GetItemByID(itemID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, item)
}

// Item godoc
// @Summary Update an existing item
// @Description Update an existing item
// @ID Update-item
// @Accept  json
// @Produce  json
// @Param id path int true "Item ID"
// @Param item body utils.Item true "Update Item"
// @Success 200 {object} utils.Item
// @Router /items/{id} [put]
func UpdateItem(c *gin.Context) {
	id := c.Param("id")
	itemID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	var updatedItem utils.Item
	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.UpdateItem(itemID, updatedItem); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, updatedItem)
}

// Item godoc
// @Summary Delete an item by ID
// @Description Delete an item by ID
// @ID Delete-item
// @Param id path int true "Item ID"
// @Success 204 "No Content"
// @Router /items/{id} [delete]
func DeleteItem(c *gin.Context) {
	id := c.Param("id")
	itemID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	if err := utils.DeleteItem(itemID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.Status(http.StatusNoContent)
}

// Item godoc
// @Summary Get all item by ID
// @Description Get all item by ID
// @ID Get-all-item
// @Accept  json
// @Produce  json
// @Success 200 {array} utils.Item
// @Router /items/ [get]
func GetAllItems(c *gin.Context) {
	items := utils.GetAllItems()
	c.JSON(http.StatusOK, items)
}
