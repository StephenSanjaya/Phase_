package main

import (
	"Phase2/week1/day2/NGC-2/config"
	"Phase2/week1/day2/NGC-2/handlers"
	"fmt"

	"net/http"
)

func main() {

	config.InitDB("root:@tcp(localhost:3306)/p2-ngc2?parseTime=true")
	defer config.DB.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/get/heroes", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetHeroes(w, r)
	})
	mux.HandleFunc("/get/villain", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetVillain(w, r)
	})

	fmt.Println("Running server on port :8081")

	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		fmt.Println("Error starting server :", err.Error())
	}
}
