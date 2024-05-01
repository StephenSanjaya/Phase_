package helpers

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

// secret key
var secretKey = "go-grpc"

/*
Notes :

  - jwt.MapClaims -> kita mau nyimpan informasi apa aja di dalam token

  - jwt.NewWithClaims -> untuk generate token
    NewWithClaims(metode encryption, data yang mau disimpan)

  - SignedString -> memasukkn secret key ke dalam token lalu mengubah tipe datanya ke bentuk string
*/
func GenerateToken(email string) string {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return signedToken
}
