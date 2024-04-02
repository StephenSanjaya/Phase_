package handlers

import (
	"Phase2/week3/day2/NGC-11-12/config"
	"Phase2/week3/day2/NGC-11-12/dto"
	"Phase2/week3/day2/NGC-11-12/entity"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterHandler(c echo.Context) error {
	user := new(dto.User)

	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid body request")
	}

	if res := config.DB.Create(&user); res.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, res.Error.Error())
	}

	return c.JSON(http.StatusCreated, "success register")
}

func LoginHandler(c echo.Context) error {
	var login dto.Login

	if err := c.Bind(&login); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	var user entity.User
	res := config.DB.Where("username = ?", login.Username).First(&user)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, res.Error)
	}
	if res.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, res.Error)
	}

	return c.JSON(http.StatusOK, "success login")
}
