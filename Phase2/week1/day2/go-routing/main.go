package main

import (
	"Phase2/week1/day2/go-routing/config"
	"Phase2/week1/day2/go-routing/handlers"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	config.InitDB("root:@tcp(localhost:3306)/hacktiv8?parseTime=true")
	defer config.DB.Close()

	router := httprouter.New()
	router.POST("/book/create", handlers.Create)
	router.GET("/books", handlers.GetAll)
	router.GET("/book/:id", handlers.GetBookByID)
	router.PUT("/book/:id", handlers.Update)
	router.DELETE("/book/:id", handlers.Delete)

	fmt.Println("Running server on port :8080")

	// running web server on local env
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error starting server :", err.Error())
	}
}
