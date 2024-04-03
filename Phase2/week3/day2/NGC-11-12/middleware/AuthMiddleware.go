package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// tokenString := c.Response().Header().Get("authorization")
		tokenString := c.Request().Header.Get("authorization")

		if tokenString == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "token not found")
		}

		parsedToken, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.NewHTTPError(http.StatusUnauthorized, "invalid argo use")
			}
			return []byte("secretkey"), nil
		})

		if parsedToken == nil || !parsedToken.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		c.Set("user_id", parsedToken.Claims.(jwt.MapClaims)["user_id"])

		return next(c)
	}
}
