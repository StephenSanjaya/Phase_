package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB(dataSourceName string) {
	var err error
	DB, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println("Error connecting to database : ", err)
	}

	err = DB.Ping()
	if err != nil {
		fmt.Println("Error ping action : ", err)
	}

	fmt.Println("Database connect!")
}
