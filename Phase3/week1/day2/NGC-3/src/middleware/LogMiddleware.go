package middleware

import (
	"Phase3/week1/day2/NGC-3/src/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func MakeLogEntry(c echo.Context) *log.Entry {
	if c == nil {
		return log.WithFields(log.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	return log.WithFields(log.Fields{
		"at":     time.Now().Format("2006-01-02 15:04:05"),
		"method": c.Request().Method,
		"uri":    c.Request().URL.String(),
		"ip":     c.Request().RemoteAddr,
		"host":   c.Request().Host,
	})
}

func MiddlewareLogging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		MakeLogEntry(c).Info("incoming request")
		return next(c)
	}
}

func ErrorHandler(err error, c echo.Context) {
	report, ok := err.(*utils.HTTPError)
	if ok {
		c.JSON(report.Code, echo.Map{
			"message": report.Message,
			"detail":  report.Detail.Error(),
		})
	} else {
		report = utils.NewHTTPError(http.StatusInternalServerError, err.Error(), err)
	}

	logError := fmt.Sprintf("http error %d - %v", report.Code, report.Detail)
	MakeLogEntry(c).Error(logError)
	// c.HTML(report.Code, report.Message.(string))
}
