package handler

import (
	"Phase2/week3/day1/latihan-slc/dto"
	"Phase2/week3/day1/latihan-slc/entity"
	"Phase2/week3/day1/latihan-slc/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	DbHandler
}

func NewCustomerHandler(dbHandler DbHandler) CustomerHandler {
	return CustomerHandler{
		DbHandler: dbHandler,
	}
}

func (ch CustomerHandler) Register(c *gin.Context) {
	var reqData dto.Customer
	if err := c.ShouldBindJSON(&reqData); err != nil {
		ErrJsonWriter(c, utils.ErrBadRequest, err.Error())
		return
	}

	data := entity.Customer{
		FirstName:   reqData.FirstName,
		LastName:    reqData.LastName,
		Email:       reqData.Email,
		Password:    reqData.Password,
		PhoneNumber: reqData.PhoneNumber,
		Address:     reqData.Address,
	}

	if hashErr := CreateHash(&data); hashErr != nil {
		ErrJsonWriter(c, *hashErr, nil)
		return
	}

	if dbErr := ch.DbHandler.InserCustomerToDb(&data); dbErr != nil {
		ErrJsonWriter(c, *dbErr, "Failed to register user into database")
		return
	}

	reqData.Password = ""
	c.JSON(http.StatusCreated, gin.H{
		"message": "Registered successfully",
		"data":    reqData,
	})
}

func (ch CustomerHandler) Login(c *gin.Context) {
	var reqData dto.Login
	if err := c.ShouldBindJSON(&reqData); err != nil {
		ErrJsonWriter(c, utils.ErrBadRequest, err.Error())
		return
	}

	dbData, dbErr := ch.DbHandler.FindUserInDb(reqData)
	if dbErr != nil {
		ErrJsonWriter(c, *dbErr, nil)
		return
	}

	if !CheckHash(dbData, reqData) {
		ErrJsonWriter(c, utils.ErrUnauthorized, "Invalid credentials")
		return
	}

	dbData.Password = ""
	c.JSON(http.StatusCreated, gin.H{
		"message": "Logged in",
		"data":    dbData,
	})
}
