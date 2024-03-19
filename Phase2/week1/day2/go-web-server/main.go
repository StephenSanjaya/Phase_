package main

import (
	"Phase2/week1/day2/go-web-server/config"
	"Phase2/week1/day2/go-web-server/handlers"
	"fmt"

	"net/http"
)

func main() {
	/*
		// Basic Web Server

		var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello world")
		}

		// set up server address
		server := http.Server{
			Addr:    "localhost:8080",
			Handler: handler,
		}

		fmt.Println("Running server on port :8080")

		// running web server on local env
		err := server.ListenAndServe()
		if err != nil {
			panic(err)
		}
	*/

	config.InitDB("root:@tcp(localhost:3306)/hacktiv8?parseTime=true")
	defer config.DB.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/create/book", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateBook(w, r)
	})
	mux.HandleFunc("/get/book", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetBooks(w, r)
	})

	fmt.Println("Running server on port :8080")

	// running web server on local env
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting server :", err.Error())
	}
}
