package handler

import (
	"Phase2/week3/day1/latihan-slc/dto"
	"Phase2/week3/day1/latihan-slc/entity"
	"Phase2/week3/day1/latihan-slc/utils"

	"gorm.io/gorm"
)

type DbHandler struct {
	*gorm.DB
}

func NewDbHandler(db *gorm.DB) DbHandler {
	return DbHandler{
		DB: db,
	}
}

func (db DbHandler) InserCustomerToDb(data *entity.Customer) *utils.ErrResponse {
	if res := db.Create(data).Error; res != nil {
		return &utils.ErrInternalServer
	}
	return nil
}

func (db DbHandler) FindUserInDb(data dto.Login) (entity.Customer, *utils.ErrResponse) {
	var user entity.Customer

	res := db.Where("email = ?", data.Email).First(&user)

	if res.RowsAffected == 0 {
		resErr := utils.ErrUnauthorized
		resErr.Details = "Invalid credentials"
		return entity.Customer{}, &resErr
	}

	if res.Error != nil {
		resErr := utils.ErrInternalServer
		return entity.Customer{}, &resErr
	}

	return user, nil
}

func (db DbHandler) FindAllProductsInDb() ([]entity.Product, *utils.ErrResponse) {
	var products []entity.Product
	if err := db.Find(&products).Error; err != nil {
		return products, &utils.ErrInternalServer
	}
	return products, nil
}
