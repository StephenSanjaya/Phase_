package config

import (
	"Phase2/week4/day1/preview-week3/entity"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("failed to load env", err)
		return nil
	}

	var dbConfig DBEnv
	err = envconfig.Process("DATABASE", &dbConfig)
	if err != nil {
		log.Fatal("failed to process env", err)
		return nil
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", dbConfig.DBHost, dbConfig.DBUsername, dbConfig.DBPassword, dbConfig.DBName, dbConfig.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database", err)
		return nil
	}

	err = db.AutoMigrate(&entity.Player{})
	if err != nil {
		log.Fatal("Failed to auto migrate database", err)
		return nil
	}

	fmt.Println("DB CONNECTED")

	return db
}
