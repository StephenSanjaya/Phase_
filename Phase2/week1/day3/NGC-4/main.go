package main

import (
	"Phase2/week1/day3/NGC-4/config"
	"Phase2/week1/day3/NGC-4/handlers"
	"fmt"

	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	config.InitDB("root:@tcp(localhost:3306)/p2-ngc2?parseTime=true")
	defer config.DB.Close()

	router := httprouter.New()
	router.GET("/reports", handlers.GetAllCriminalReports)
	router.GET("/reports/:id", handlers.GetCriminalReportById)
	router.POST("/reports", handlers.CreateReport)
	router.PUT("/reports/:id", handlers.UpdateReportById)
	router.DELETE("/reports/:id", handlers.DeleteReportById)

	fmt.Println("Running server on port :8081")

	err := http.ListenAndServe(":8081", router)
	if err != nil {
		fmt.Println("Error starting server :", err.Error())
	}
}
