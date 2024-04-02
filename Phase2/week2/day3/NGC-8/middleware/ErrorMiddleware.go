package middleware

import (
	"Phase2/week2/day3/NGC-8/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorMiddleware(c *gin.Context) {
	c.Next()
	// fmt.Println(reflect.TypeOf(c.Errors.Last()))

	// for _, err := range c.Errors {
	// 	logger(err.Error())
	// }

	switch c.Errors.Last().Error() {
	case "BadRequest":
		c.JSON(http.StatusBadRequest, models.ErrorContract{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: "Invalid body input",
			ErrorDetail:  "",
		})
	default:
		// c.JSON(500, gin.H{"error": ErrNotFound.Error()})
	}

	// c.JSON(http.StatusInternalServerError, "")
}
