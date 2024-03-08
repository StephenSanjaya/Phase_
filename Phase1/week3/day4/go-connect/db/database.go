package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/go_db")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to DB: ", err.Error())
		return nil, err
	} 

	fmt.Println("Connected !")
	return db, nil
}

func ExecQuery(db *sql.DB, query string, value string) error {
	_, err := db.Exec(query, value)
	if err != nil {
		return err
	}

	return nil
}

func ExecQueryCRUD(db *sql.DB, query string) error {
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}