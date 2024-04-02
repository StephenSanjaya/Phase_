package main

import (
	"Phase2/week2/day4/go-swagger/router"
)

func main() {

	// Menginisialisasi router
	r := router.SetupRouter()

	// Menentukan port yang akan digunakan
	port := ":8080" // Ganti dengan port yang sesuai

	// Menjalankan server
	if err := r.Run(port); err != nil {
		panic(err)
	}
}
