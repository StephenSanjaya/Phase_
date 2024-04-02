package handler

import (
	"Phase2/week3/day1/latihan-slc/dto"
	"Phase2/week3/day1/latihan-slc/entity"
	"Phase2/week3/day1/latihan-slc/utils"

	"golang.org/x/crypto/bcrypt"
)

func CreateHash(data *entity.Customer) *utils.ErrResponse {
	hashed, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		hashErr := utils.ErrInternalServer
		hashErr.Details = err.Error()
		return &hashErr
	}

	data.Password = string(hashed)
	return nil
}

func CheckHash(dbData entity.Customer, data dto.Login) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(dbData.Password), []byte(data.Password)); err != nil {
		return false
	}
	return true
}
