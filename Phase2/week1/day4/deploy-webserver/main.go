package main

import (
	"Phase2/week1/day4/deploy-webserver/config"
	"Phase2/week1/day4/deploy-webserver/handlers"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
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

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var dbConfig config.DBEnv
	err = envconfig.Process("DATABASE", &dbConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	source := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName)
	config.InitDB(source)
	defer config.DB.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/create/book", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateBook(w, r)
	})
	mux.HandleFunc("/get/books", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetBooks(w, r)
	})

	fmt.Println("Running server on port :8080")

	// running web server on local env
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting server :", err.Error())
	}
}
