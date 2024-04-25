package main

import (
	"Phase3/week1/day2/go-dependency-injection-04/config"
	"context"
)

func main() {
	db := config.ConnectDB()
	defer db.Disconnect(context.Background())
}
