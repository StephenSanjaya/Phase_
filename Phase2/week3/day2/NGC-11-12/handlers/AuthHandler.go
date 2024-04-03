package handlers

import (
	"Phase2/week3/day2/NGC-11-12/config"
	"Phase2/week3/day2/NGC-11-12/dto"
	"Phase2/week3/day2/NGC-11-12/entity"
	"errors"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Register godoc
// @Summary Create user
// @Description Register new user
// @ID Create-user
// @Param register body dto.User true "User to register"
// @Accept  json
// @Produce  json
// @Success      201              {string}  string    "success register"
// @Failure      400              {string}  string    "bad request"
// @Failure      404              {string}  string    "not found"
// @Failure      500              {string}  string    "internal server error"
// @Router /register [post]
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

// Register godoc
// @Summary Create user
// @Description Register new user
// @ID Login-user
// @Param login body dto.Login true "User to login"
// @Accept  json
// @Produce  json
// @Success      200              {string}  string    "success login"
// @failure      400              {string}  string    "bad request"
// @failure      404              {string}  string    "not found"
// @failure      500              {string}  string    "internal server error"
// @Router /login [post]
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

	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte("secretkey"))
	if err != nil {
		cusError := fmt.Sprintf("failed create token: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, cusError)
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "success login",
		"jwt":     tokenString,
	})
}
