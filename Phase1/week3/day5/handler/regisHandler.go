package handler

import (
	"Phase1/week3/day5/config"
	"Phase1/week3/day5/entity"

	"golang.org/x/crypto/bcrypt"
)

func Register(user entity.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	_, err = config.DB.Exec("INSERT INTO users(email, fullname, password, dob) VALUES (?,?,?,?) ", user.Email, user.Fullname, hashedPassword, user.Dob)

	if err != nil {
		return err
	}
	return nil
}
