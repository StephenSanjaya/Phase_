package main

import (
	handler "Phase2/week3/day3/go-third-5/Handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/country", handler.HandlerCountry)
	r.Run(":8081")
}
