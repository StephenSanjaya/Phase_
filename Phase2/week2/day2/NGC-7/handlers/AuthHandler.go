package handlers

import (
	"Phase2/week2/day2/NGC-7/config"
	"Phase2/week2/day2/NGC-7/models"
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
		c.JSON(http.StatusBadRequest, models.ErrorMessage{
			Status:  http.StatusBadRequest,
			Message: "invalid body input",
		})
		return
	}

	// validasi request body
	// err = helper.ValidateUser(store)
	// if err != nil {
	// 	json.NewEncoder(w).Encode(models.ErrorMessage{
	// 		Status:  http.StatusBadRequest,
	// 		Message: err.Error(),
	// 	})
	// 	return
	// }

	stmt, err := config.Db.Prepare("INSERT INTO Stores (email, password, store_name, store_type) VALUES (?,?,?,?)")
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorMessage{
			Status:  http.StatusInternalServerError,
			Message: "failed to prepare query: " + err.Error(),
		})
		return
	}
	defer stmt.Close()

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(store.Password), bcrypt.DefaultCost)

	_, err = stmt.Exec(store.Email, hashedPassword, store.StoreName, store.StoreType)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorMessage{
			Status:  http.StatusInternalServerError,
			Message: "failed to register: " + err.Error(),
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
		c.JSON(http.StatusBadRequest, models.ErrorMessage{
			Status:  http.StatusBadRequest,
			Message: "invalid body input",
		})
		return
	}

	stmt, err := config.Db.Prepare("SELECT email, password, store_name, store_type FROM Stores WHERE email = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorMessage{
			Status:  http.StatusInternalServerError,
			Message: "failed to prepare query: " + err.Error(),
		})
		return
	}
	defer stmt.Close()

	var store models.Store
	var hashedPassword *string

	rows := stmt.QueryRow(userLogin.Email).Scan(&store.Email, &hashedPassword, &store.StoreName, &store.StoreType)
	if rows != nil {
		c.JSON(http.StatusNotFound, models.ErrorMessage{
			Status:  http.StatusNotFound,
			Message: "email not found: " + rows.Error(),
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(*hashedPassword), []byte(userLogin.Password))
	if err != nil {
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
		c.JSON(http.StatusNotFound, models.ErrorMessage{
			Status:  http.StatusInternalServerError,
			Message: "failed create token: " + err.Error(),
		})
		return
	}

	var SuccessMessage = models.SuccessMessage{
		Status:  http.StatusAccepted,
		Message: "Login Success",
	}

	//create cookie - harusnya gaperlu, bisa passing lewat context
	// http.SetCookie(w, &http.Cookie{
	// 	Name:     "role",
	// 	Path:     "/",
	// 	Value:    user.Role,
	// 	HttpOnly: true,
	// })

	c.Writer.Write([]byte(tokenString))
	c.JSON(http.StatusAccepted, SuccessMessage)
}
