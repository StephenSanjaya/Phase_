package main

import (
	"fmt"
	"Phase1/week3/day4/go-connect/db"
)

type Test struct {
	id	int
	name string
}

func main() {
	database, err := db.ConnectDB()
	if err != nil {
		fmt.Println(err.Error())
		return 
	}

	defer database.Close()

	err = db.ExecQueryCRUD(database, "CREATE TABLE IF NOT EXISTS test (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(50))")
	if err != nil {
		fmt.Println("Failed to create table : ", err.Error())
		return
	}

	fmt.Println("Success to create table!")

	err = db.ExecQuery(database, "INSERT INTO test (name) VALUES (?)", "Budi")
	if err != nil {
		fmt.Println("Failed to insert data : ", err.Error())
		return
	}

	fmt.Println("Success to insert data!")


	rows, err := database.Query("SELECT * FROM test")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var tests []Test

		err = rows.Scan(&id, &name)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		tests = append(tests, Test{id, name})
		fmt.Println(tests)

		// fmt.Println("ID: ", id)
		// fmt.Println("Name: ", id)
	}

}