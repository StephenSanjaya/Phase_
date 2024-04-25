package controller

import (
	"Phase3/week1/day2/NGC-3/src/model"
	"Phase3/week1/day2/NGC-3/src/repository"
	"Phase3/week1/day2/NGC-3/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageController struct {
	Repository repository.MessageRepositoryI
	Validate   *validator.Validate
}

func NewMessageController(mr repository.MessageRepositoryI, validate *validator.Validate) MessageControllerI {
	return &MessageController{Repository: mr, Validate: validate}
}

func (mc *MessageController) CreateMessage(c echo.Context) error {
	msg := new(model.Message)

	if err := c.Bind(&msg); err != nil {
		return utils.NewHTTPError(http.StatusBadRequest, "invalid body request", err)
	}

	if err := mc.Validate.Struct(msg); err != nil {
		return utils.NewHTTPError(http.StatusBadRequest, "invalid body request", err)
	}

	msg.UserReceiver.Name = "me"
	msg.UserReceiver.PhoneNumber = "08123456789"

	err := mc.Repository.Save(msg)
	if err != nil {
		return utils.NewHTTPError(err.Code, err.Message, err.Detail)
	}

	return c.JSON(http.StatusCreated, gin.H{
		"message": "success create message",
	})
}

func (mc *MessageController) FindMessageByID(c echo.Context) error {
	message_id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return utils.NewHTTPError(http.StatusBadRequest, "failed to parse id", err)
	}

	msg, errEcho := mc.Repository.FindByID(message_id)
	if errEcho != nil {
		return utils.NewHTTPError(errEcho.Code, errEcho.Message, errEcho.Detail)
	}

	return c.JSON(http.StatusOK, gin.H{
		"message": "success find message by id",
		"data":    msg,
	})
}

func (mc *MessageController) FindAllMessageBySender(c echo.Context) error {
	name := c.Param("name")

	msgs, err := mc.Repository.FindAllMessageBySender(name)
	if err != nil {
		return utils.NewHTTPError(err.Code, err.Message, err.Detail)
	}

	return c.JSON(http.StatusOK, gin.H{
		"message": "success get all message by sender",
		"data":    msgs,
	})
}
