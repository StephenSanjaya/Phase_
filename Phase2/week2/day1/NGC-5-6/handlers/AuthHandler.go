package handlers

import (
	"Phase2/week2/day1/NGC-5-6/config"
	"Phase2/week2/day1/NGC-5-6/helper"
	"Phase2/week2/day1/NGC-5-6/models"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Status:  http.StatusBadRequest,
			Message: "invalid body input",
		})
		return
	}

	err = helper.ValidateUser(user)
	if err != nil {
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	stmt, err := config.Db.Prepare("INSERT INTO Users (email, password, full_name, age, occupation, role) VALUES (?,?,?,?,?,?)")
	if err != nil {
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Status:  http.StatusInternalServerError,
			Message: "failed to prepare query: " + err.Error(),
		})
		return
	}
	defer stmt.Close()

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	_, err = stmt.Exec(user.Email, hashedPassword, user.FullName, user.Age, user.Occupation, user.Role)
	if err != nil {
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Status:  http.StatusInternalServerError,
			Message: "failed to register: " + err.Error(),
		})
		return
	}

	var SuccessMessage = models.SuccessMessage{
		Status:  http.StatusCreated,
		Message: "Success register new user",
	}

	log.Default().Println("HTTP request sent to POST /register")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(SuccessMessage)

}

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	var userLogin models.Login
	err := json.NewDecoder(r.Body).Decode(&userLogin)
	if err != nil {
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Status:  http.StatusBadRequest,
			Message: "invalid body input",
		})
		return
	}

	stmt, err := config.Db.Prepare("SELECT user_id, email, password, full_name, age, occupation, role FROM Users WHERE email = ?")
	if err != nil {
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Status:  http.StatusInternalServerError,
			Message: "failed to prepare query: " + err.Error(),
		})
		return
	}
	defer stmt.Close()

	var user models.User
	var hashedPassword *string

	rows := stmt.QueryRow(userLogin.Email).Scan(&user.UserID, &user.Email, &hashedPassword, &user.FullName, &user.Age, &user.Occupation, &user.Role)
	if rows != nil {
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Status:  http.StatusNotFound,
			Message: "email not found: " + rows.Error(),
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(*hashedPassword), []byte(userLogin.Password))
	if err != nil {
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Status:  http.StatusNotFound,
			Message: "password not match: " + err.Error(),
		})
		return
	}

	//creat jwt
	claims := jwt.MapClaims{
		"user_id":    user.UserID,
		"email":      user.Email,
		"username":   user.FullName,
		"age":        user.Age,
		"occupation": user.Occupation,
		"role":       user.Role,
		"exp":        time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret_token := []byte("mysecretKey")

	tokenString, err := token.SignedString(secret_token)
	if err != nil {
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Status:  http.StatusInternalServerError,
			Message: "failed create token: " + err.Error(),
		})
		return
	}

	var SuccessMessage = models.SuccessMessage{
		Status:  http.StatusAccepted,
		Message: "Login Success",
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "role",
		Path:     "/",
		Value:    user.Role,
		HttpOnly: true,
	})

	log.Default().Println("HTTP request sent to POST /login")

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(tokenString))
	json.NewEncoder(w).Encode(SuccessMessage)
}
