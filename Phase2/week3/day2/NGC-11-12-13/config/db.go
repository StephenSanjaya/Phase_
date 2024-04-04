package config

import (
	"Phase2/week3/day2/NGC-11-12-13/entity"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase(dataSourceName string) {
	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed connect db: ", err)
		return
	}

	err = db.SetupJoinTable(&entity.User{}, "Products", &entity.Transaction{})
	if err != nil {
		log.Fatal("Failed to setup join table: ", err)
		return
	}

	err = db.AutoMigrate(&entity.Product{}, &entity.User{}, &entity.Transaction{})
	if err != nil {
		log.Fatal("Failed to migrate db: ", err)
	}

	DB = db
	fmt.Println("Database connected")
}
