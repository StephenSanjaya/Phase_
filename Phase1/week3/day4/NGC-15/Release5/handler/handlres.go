package handler

import (
	"Phase1/week3/day4/NGC-15/Release5/entity"
	"context"
	"database/sql"
	"fmt"
	"time"
)

func InsertNewEmployee(db *sql.DB, employee entity.Employee) error {
	_, err := db.Exec("INSERT INTO Employees (first_name, last_name, position) VALUES (?,?,?)", employee.First_name, employee.Last_name, employee.Position)
	if err != nil {
		fmt.Println("Error executing query: ", err.Error())
		return err
	}

	return nil
}

func GetAllEmployees(db *sql.DB) (listEmployees []entity.Employee,err error)  {
	var e = entity.Employee{}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, "SELECT first_name, last_name, position FROM Employees")
	if err != nil {
		fmt.Println("Error executing query:", err.Error())
		return listEmployees, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&e.First_name, &e.Last_name, &e.Position)
		if err != nil {
			fmt.Println(err.Error())
			return listEmployees, err
		}
		listEmployees = append(listEmployees, e)
	}

	return listEmployees, nil
}

func InsertNewMenu(db *sql.DB, menu entity.MenuItem) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx, "INSERT INTO MenuItems (name, description, price, category) VALUES (?,?,?,?)", menu.Name, menu.Description, menu.Price, menu.Category)
	if err != nil {
		fmt.Println("Error executing query: ", err.Error())
		return err
	}

	return nil
}

func GetAllMenuItems(db *sql.DB) (listMenuItems []entity.MenuItem,err error) {
	var mi = entity.MenuItem{}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, "SELECT name, description, price, category FROM MenuItems")
	if err != nil {
		fmt.Println("Error executing query:", err.Error())
		return listMenuItems, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&mi.Name, &mi.Description, &mi.Price, &mi.Category)
		if err != nil {
			fmt.Println(err.Error())
			return listMenuItems, err
		}
		listMenuItems = append(listMenuItems, mi)
	}

	return listMenuItems, nil
}

func CreateNewOrder(db *sql.DB, order entity.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx, "INSERT INTO Orders (table_number, employee_id, order_date, status) VALUES (?,?,?,?)", order.Table_number, order.Employee_id, order.Order_date, order.Status)
	if err != nil {
		fmt.Println("Error executing query: ", err.Error())
		return err
	}

	return nil
}

func GetCurrentOrderID(db *sql.DB, tableNum int) (orderID int, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.QueryRowContext(ctx, "SELECT order_id FROM Orders WHERE table_number = ?", tableNum).Scan(&orderID)
	if err != nil {
		fmt.Println("Error executing query:", err.Error())
		return -1, err
	}

	return orderID, nil
}

func GetItemPrice(db *sql.DB, itemID int) (price float64, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.QueryRowContext(ctx, "SELECT price FROM MenuItems WHERE item_id = ?", itemID).Scan(&price)
	if err != nil {
		fmt.Println("Error executing query:", err.Error())
		return -1, err
	}

	return price, nil
}

func CreateNewOrderItem(db *sql.DB, orderitem entity.OrderItems) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx, "INSERT INTO OrderItems (order_id, item_id, quantity, subtotal) VALUES (?,?,?,?)", orderitem.Order_id, orderitem.Item_id, orderitem.Quantity, orderitem.Subtotal)
	if err != nil {
		fmt.Println("Error executing query: ", err.Error())
		return err
	}

	return nil
}

func GetAllOrdersDetail(db *sql.DB) (orderdetails []entity.OrderDetails, err error) {
	var od = entity.OrderDetails{}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, "SELECT ot.order_id, o.table_number, SUM(ot.subtotal) AS Total, GROUP_CONCAT(mt.name SEPARATOR ', ') FROM OrderItems ot JOIN MenuItems mt ON ot.item_id = mt.item_id JOIN Orders o ON ot.order_id = o.order_id WHERE o.status = ? GROUP BY ot.order_id", "Placed")
	if err != nil {
		fmt.Println("Error executing query:", err.Error())
		return orderdetails, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&od.OrderID, &od.TableNumber, &od.TotalPrice, &od.Name)
		if err != nil {
			fmt.Println(err.Error())
			return orderdetails, err
		}
		orderdetails = append(orderdetails, od)
	}

	return orderdetails, nil
}

func GetTotalAmountFromOrderID(db *sql.DB, orderID int) (totalamount float64, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.QueryRowContext(ctx, "SELECT SUM(ot.subtotal) AS Total FROM OrderItems ot JOIN MenuItems mt ON ot.item_id = mt.item_id JOIN Orders o ON ot.order_id = o.order_id WHERE o.status = ? AND ot.order_id = ? GROUP BY ot.order_id", "Placed", orderID).Scan(&totalamount)
	if err != nil {
		fmt.Println("Error executing query:", err.Error())
		return -1, err
	}

	return totalamount, nil
}

func InsertNewPaymentOrder(db *sql.DB, payment entity.Payment) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx, "INSERT INTO Payments (order_id, payment_date, payment_method, total_amount) VALUES (?,?,?,?)", payment.OrderID, payment.PaymentDate, payment.PaymentMethod, payment.TotalAmount)
	if err != nil {
		fmt.Println("Error executing query: ", err.Error())
		return err
	}

	return nil
}

func ChangeStatusOrder(db *sql.DB, orderID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.ExecContext(ctx, "UPDATE Orders SET Status = ? WHERE order_id = ?", "Completed", orderID)
	if err != nil {
		fmt.Println("Error executing query: ", err.Error())
		return err
	}

	return nil
}