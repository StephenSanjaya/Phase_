package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			// http.Error(c.Writer, "Unauthorized", http.StatusUnauthorized)
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
			return
		}

		secret_token := []byte(os.Getenv("JWT"))
		parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid algorithm use")
				return nil, fmt.Errorf("invalid algorithm use")
			}
			return secret_token, nil
		})

		if parsedToken == nil || !parsedToken.Valid {
			fmt.Println("error while decode token", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid token")
			return
		}

		c.Next()
	}
}
