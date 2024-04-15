package middleware

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("authorization")

		if tokenString == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "token not found")
		}

		parsedToken, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.NewHTTPError(http.StatusUnauthorized, "invalid algo use")
			}
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if parsedToken == nil || !parsedToken.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		if float64(time.Now().Unix()) > parsedToken.Claims.(jwt.MapClaims)["exp"].(float64) {
			return echo.NewHTTPError(http.StatusUnauthorized, errors.New("token expired"))
		}

		c.Set("email", parsedToken.Claims.(jwt.MapClaims)["email"])

		return next(c)
	}
}
