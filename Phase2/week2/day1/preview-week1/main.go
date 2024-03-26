package main

import (
	"Phase2/week2/day1/preview-week1/config"
	"Phase2/week2/day1/preview-week1/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/kelseyhightower/envconfig"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("err loading env")
		return
	}

	var dbConfig config.DBEnv
	err = envconfig.Process("DATABASE", &dbConfig)
	if err != nil {
		log.Fatal("err process env")
		return
	}

	source := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", dbConfig.DBUsername, dbConfig.DBPassword, dbConfig.DBHost, dbConfig.DBPort, dbConfig.DBName)

	config.GetConnection(source)
	defer config.Db.Close()

	router := httprouter.New()
	router.GET("/branches", handlers.GetAllBranches)
	router.GET("/branches/:id", handlers.GetBranchById)
	router.POST("/branches", handlers.InsertNewBranches)
	router.PUT("/branches/:id", handlers.UpdateBranch)
	router.DELETE("/branches/:id", handlers.DeleteBranch)

	err = http.ListenAndServe(":8081", router)
	if err != nil {
		fmt.Println("failed to listen to port 8081, " + err.Error())
		return
	}

}
