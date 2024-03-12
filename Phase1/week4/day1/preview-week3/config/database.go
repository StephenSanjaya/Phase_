package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/p1-pw3")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return db, nil
}