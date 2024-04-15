package handler

import (
	"net/http"
	"os"
	"time"

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

type User struct {
	ID       int
	Email    string
	Password string
}

func (as *AuthService) LoginHandler(c echo.Context) error {
	user := new(User)
	user.ID = 1
	user.Email = "admin@gmail.com"
	user.Password = "admin123"

	// user := new(User)
	// user.ID = 2
	// user.Email = "user@gmail.com"
	// user.Password = "user123"

	accessToken, refreshToken, err := CreateToken(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"access-token":  accessToken,
		"refresh-token": refreshToken,
	})
}

func (as *AuthService) RefreshTokenHandler(c echo.Context) error {
	user := new(User)
	user.ID = 1
	user.Email = "admin@gmail.com"
	user.Password = "admin123"

	refreshToken := c.Request().Header.Get("authorization")

	parsedToken, err := jwt.Parse(refreshToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, echo.NewHTTPError(http.StatusUnauthorized, "invalid argo use")
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if parsedToken == nil || !parsedToken.Valid {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	accessToken, refreshToken, err := CreateToken(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"access-token":  accessToken,
		"refresh-token": refreshToken,
	})
}

func CreateToken(user *User) (string, string, error) {
	claims := jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Minute * 5).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", "", err
	}

	claims2 := jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims2)

	refreshTokenString, err := refreshToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", "", err
	}

	return tokenString, refreshTokenString, nil
}
