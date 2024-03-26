package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var secretKey = []byte("12345")

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	//NewWithClaims
	claims := jwt.MapClaims{
		"username": "tugas",
		"role":     "admin",
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		fmt.Println("Failed create token")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(tokenString))

}

func AuthMiddleware(next http.Handler) http.Handler {
	//Parse
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if path == "/login" {
			next.ServeHTTP(w, r)
			return
		}

		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			http.Error(w, "token not found", http.StatusUnauthorized)
			return
		}

		parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("algoritma tidak valid")
			}
			return secretKey, nil
		})

		if parsedToken.Valid {
			fmt.Println(parsedToken)
		}

		if parsedToken == nil || !parsedToken.Valid {
			fmt.Println("error while decode token", err)
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)

	})
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("halaman ini setelah login"))
}
func main() {

	r := mux.NewRouter()
	r.HandleFunc("/login", LoginHandler).Methods("POST")

	r.Use(AuthMiddleware)
	r.HandleFunc("/home", Home).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
