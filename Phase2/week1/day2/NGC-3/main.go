package main

import (
	"Phase2/week1/day2/NGC-3/config"
	"Phase2/week1/day2/NGC-3/handlers"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	config.InitDB("root:@tcp(localhost:3306)/p2-ngc2?parseTime=true")
	defer config.DB.Close()

	router := httprouter.New()
	router.GET("/inventories", handlers.GetAllInventories)
	router.GET("/inventories/:id", handlers.GetInventoryByID)
	router.POST("/inventories", handlers.CreateItem)
	router.PUT("/inventories/:id", handlers.UpdateInventoryById)
	router.DELETE("/inventories/:id", handlers.DeleteInventoryById)

	fmt.Println("Running server on port :8081")

	// running web server on local env
	err := http.ListenAndServe(":8081", router)
	if err != nil {
		fmt.Println("Error starting server :", err.Error())
	}
}
