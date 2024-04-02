package config

import (
	"Phase2/week2/day3/NGC-8/models"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func GetConnection() {
	var err error

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("err loading env")
		return
	}

	var dbConfig DBEnv
	err = envconfig.Process("DATABASE", &dbConfig)
	if err != nil {
		log.Fatal("err process env")
		return
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", dbConfig.DBHost, dbConfig.DBUsername, dbConfig.DBPassword, dbConfig.DBName, dbConfig.DBPort)

	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed connect db", err)
		return
	}

	Db.AutoMigrate(&models.Store{}, &models.Product{})

	fmt.Println("DB CONNECTED")
}
