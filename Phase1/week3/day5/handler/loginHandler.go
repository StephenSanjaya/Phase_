package handler

import (
	"Phase1/week3/day5/config"
	"Phase1/week3/day5/entity"

	"golang.org/x/crypto/bcrypt"
)

func Login(fullname string, password string) (entity.User, bool) {
	var user entity.User
	row := config.DB.QueryRow("SELECT id, fullname, password FROM users WHERE fullname = ?", fullname)
	err := row.Scan(&user.ID, &user.Fullname, &user.Password)
	if err != nil {
		return user, false
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, false
	}

	return user, true
}
