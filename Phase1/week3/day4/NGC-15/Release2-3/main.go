package main

import (
	"Phase1/week3/day4/NGC-15/db"
	"context"
	"fmt"
	"os"
	"strings"
	"time"
)

func handlePanic()  {
	if r := recover(); r != nil {
		fmt.Println("Recovered form panic: ", r)
	}
}

func main() {

	defer handlePanic()

	args := os.Args[1:]
	if len(args) < 1{
		panic("No SQL Provided!")
	}

	sqlFile := args[0]
	sqlCommands, err := os.ReadFile(sqlFile)
	if err != nil {
		panic(err)
	}

	database, err := db.ConnectDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer database.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	arrSqlCommands := strings.Split(string(sqlCommands), ";")
	for i := 0; i < len(arrSqlCommands)-1; i++ {
		_, err = database.ExecContext(ctx, arrSqlCommands[i])
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Database Migration sucessfull!")

}