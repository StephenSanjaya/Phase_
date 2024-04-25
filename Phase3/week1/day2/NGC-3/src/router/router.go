package router

import (
	"Phase3/week1/day2/NGC-3/src/controller"

	"github.com/labstack/echo/v4"
)

func SetupRouter(e *echo.Echo, msgController controller.MessageControllerI) {

	e.POST("/message", msgController.CreateMessage)
	e.GET("/message/:id", msgController.FindMessageByID)
	e.GET("/message/sender/:name", msgController.FindAllMessageBySender)

}
