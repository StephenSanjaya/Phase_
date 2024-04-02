package config

import (
	"Phase2/week3/day1/latihan-slc/entity"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb() *gorm.DB {
	dsn := os.Getenv("DB")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err1 := db.SetupJoinTable(&entity.Order{}, "Products", &entity.OrderItem{})
	if err1 != nil {
		log.Fatal(err1)
	}
	err2 := db.AutoMigrate(&entity.Customer{}, &entity.Product{}, &entity.Payment{}, &entity.Order{}, &entity.OrderItem{})
	if err2 != nil {
		log.Fatal(err2)
	}
	return db
}
