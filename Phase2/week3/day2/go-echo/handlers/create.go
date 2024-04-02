package handlers

import (
	"Phase2/week3/day2/go-echo/config"
	"Phase2/week3/day2/go-echo/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateBook(c echo.Context) error {
	book := new(model.Book)

	// binding payload into model
	if err := c.Bind(book); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid payload")
	}

	// create data into database
	result := config.DB.Create(&book)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error.Error())
	}

	// create response
	return c.JSON(http.StatusCreated, book)
}
