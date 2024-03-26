package main

import (
	"Phase2/week2/day1/NGC-5-6/config"
	"Phase2/week2/day1/NGC-5-6/handlers"
	"Phase2/week2/day1/NGC-5-6/middleware"
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
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)

	router.GET("/recipes", middleware.AuthMiddleware(handlers.GetAllRecipes))
	router.GET("/recipes/:id", middleware.AuthMiddleware(handlers.GetRecipeById))
	router.POST("/recipes", middleware.AuthMiddleware(handlers.InsertNewRecipe))
	router.PUT("/recipes/:id", middleware.AuthMiddleware(handlers.UpdateRecipe))
	router.DELETE("/recipes/:id", middleware.AuthMiddleware(handlers.DeleteRecipe))

	err = http.ListenAndServe(":8081", router)
	if err != nil {
		fmt.Println("failed to listen to port 8081, " + err.Error())
		return
	}

}
