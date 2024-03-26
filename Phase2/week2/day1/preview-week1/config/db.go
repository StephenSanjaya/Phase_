package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func GetConnection(source string) {
	var err error
	Db, err = sql.Open("mysql", source)
	if err != nil {
		fmt.Println("failed connect db", err)
		return
	}

	err = Db.Ping()
	if err != nil {
		fmt.Println("failed ping db", err)
		return
	}

	fmt.Println("DB CONNECTED")
}
