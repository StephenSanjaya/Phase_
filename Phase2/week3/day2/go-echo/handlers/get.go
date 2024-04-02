package handlers

import (
	"Phase2/week3/day2/go-echo/config"
	"Phase2/week3/day2/go-echo/model"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetBooks(c echo.Context) error {
	var books []model.Book

	result := config.DB.Find(&books)

	// handle error 404
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, "Book not found")
	}

	// handle error 500
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error.Error())
	}

	return c.JSON(http.StatusOK, books)
}

func GetBookById(c echo.Context) error {
	var book model.Book
	id := c.Param("id")

	result := config.DB.First(&book, id)

	// handle error 404
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, "Book not found")
	}

	// handle error 500
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error.Error())
	}

	return c.JSON(http.StatusOK, book)
}
