package main

import (
	"Phase2/week1/day4/deploy-ngc4/config"
	"Phase2/week1/day4/deploy-ngc4/handlers"
	"fmt"
	"log"
	"os"

	"net/http"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/kelseyhightower/envconfig"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	var dbConfig config.DBEnv
	err = envconfig.Process("DATABASE", &dbConfig)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	source := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName)

	config.InitDB(source)
	defer config.DB.Close()

	router := httprouter.New()
	router.GET("/reports", handlers.GetAllCriminalReports)
	router.GET("/reports/:id", handlers.GetCriminalReportById)
	router.POST("/reports", handlers.CreateReport)
	router.PUT("/reports/:id", handlers.UpdateReportById)
	router.DELETE("/reports/:id", handlers.DeleteReportById)

	fmt.Println("Running server on port :8081")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Println("Error starting server :", err.Error())
	}
}
