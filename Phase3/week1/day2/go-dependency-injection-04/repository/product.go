package repository

import (
	"Phase3/week1/day2/go-dependency-injection-04/config/utils"
	"Phase3/week1/day2/go-dependency-injection-04/model"
)

type Product interface {
	Create(data *model.Product) *utils.ErrResponse
	// GetById(string) error
	// GetAll() error
	// Update() error
	// Delete() error
}
