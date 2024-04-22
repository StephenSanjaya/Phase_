package handler

import (
	"Phase3/week1/day1/preview-phase2/dto"
	"Phase3/week1/day1/preview-phase2/entity"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{db: db}
}

func (as *AuthService) RegisterHandler(c echo.Context) error {
	user := new(entity.User)

	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid body request")
	}

	if res := as.db.Create(&user); res.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, res.Error.Error())
	}

	user.Password = ""

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "success register",
		"user":    user,
	})
}

func (as *AuthService) LoginHandler(c echo.Context) error {
	var login dto.Login

	if err := c.Bind(&login); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	var user entity.User
	res := as.db.Where("username = ?", login.Email).First(&user)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, res.Error)
	}
	if res.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, res.Error)
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		cusError := fmt.Sprintf("failed create token: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, cusError)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success login",
		"jwt":     tokenString,
	})
}
