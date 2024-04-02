package handlers

import (
	"Phase2/week2/day3/NGC-8/config"
	"Phase2/week2/day3/NGC-8/models"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")

	var store models.Store
	if err := c.ShouldBindJSON(&store); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorMessage{
			Status:  http.StatusBadRequest,
			Message: "Invalid body input: " + err.Error(),
		})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(store.Password), bcrypt.DefaultCost)
	store.Password = string(hashedPassword)

	if res := config.Db.Create(&store); res.Error != nil {

		c.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorMessage{
			Status:  http.StatusInternalServerError,
			Message: "Failed to register new user: " + res.Error.Error(),
		})
		return
	}

	var SuccessMessage = models.SuccessMessage{
		Status:  http.StatusCreated,
		Message: "Success register new store",
	}

	c.JSON(http.StatusCreated, SuccessMessage)

}

func LoginHandler(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")

	var userLogin models.Login
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		custErr := errors.New("BadRequest") //pke logrus
		c.Error(custErr)
		fmt.Println(err)
		// c.Error(err)
		// c.JSON(http.StatusBadRequest, models.ErrorMessage{
		// 	Status:  http.StatusBadRequest,
		// 	Message: "invalid body input: " + err.Error(),
		// })
		return
	}

	var store models.Store
	if res := config.Db.Where("email = ?", userLogin.Email).First(&store); res.Error != nil {
		c.Error(res.Error)
		c.AbortWithStatusJSON(http.StatusNotFound, models.ErrorMessage{
			Status:  http.StatusNotFound,
			Message: "Email not found" + res.Error.Error(),
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(store.Password), []byte(userLogin.Password))
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusNotFound, models.ErrorMessage{
			Status:  http.StatusNotFound,
			Message: "password not match: " + err.Error(),
		})
		return
	}

	//creat jwt
	claims := jwt.MapClaims{
		"email":      store.Email,
		"store_name": store.StoreName,
		"store_type": store.StoreType,
		"exp":        time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret_token := []byte(os.Getenv("JWT"))

	tokenString, err := token.SignedString(secret_token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorMessage{
			Status:  http.StatusInternalServerError,
			Message: "failed create token: " + err.Error(),
		})
		return
	}

	var SuccessMessage = models.SuccessMessage{
		Status:  http.StatusAccepted,
		Message: "Login Success",
	}

	c.Writer.Write([]byte(tokenString))
	c.JSON(http.StatusAccepted, SuccessMessage)
}
