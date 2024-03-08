package main

import (
	"Phase1/week3/day4/NGC-15/db"
	"context"
	"fmt"
	"time"
)

func main() {
	database, err := db.ConnectDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer database.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	fmt.Println("=========================================")

	rows, err := database.QueryContext(ctx, "SELECT o.order_id, e.first_name, e.position, o.table_number, o.order_date, o.status FROM Orders o JOIN Employees e ON o.employee_id = e.employee_id")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var order_id, table_number int
		var first_name, position, order_date, status string
		err := rows.Scan(&order_id, &first_name, &position, &table_number, &order_date, &status)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Order ID: \t", order_id)
		fmt.Println("Employee Name: \t", first_name)
		fmt.Println("Position: \t", position)
		fmt.Println("Table Number: \t", table_number)
		fmt.Println("Order Date: \t", order_date)
		fmt.Println("Status: \t", status)
		fmt.Println()
	}

	fmt.Println("=========================================")

	rows, err = database.QueryContext(ctx, "SELECT ot.order_id, SUM(mt.price) * ot.quantity AS Total, GROUP_CONCAT(mt.name SEPARATOR ', ') FROM OrderItems ot JOIN MenuItems mt ON ot.item_id = mt.item_id GROUP BY ot.order_id")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var order_id int
		var total float64
		var names string
		err := rows.Scan(&order_id, &total, &names)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Order ID: \t", order_id)
		fmt.Println("Names: \t\t", names)
		fmt.Println("Total Price: \t", total)
		fmt.Println()
	}
}