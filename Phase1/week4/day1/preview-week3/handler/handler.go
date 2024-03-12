package handler

import (
	"Phase1/week4/day1/preview-week3/entity"
	"context"
	"database/sql"
	"fmt"
	"time"
)

func GetListAllBooks(db *sql.DB) (books []entity.Books, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	var b = entity.Books{}

	rows, err := db.QueryContext(ctx, "SELECT b.book_title FROM Books b JOIN Authors a ON b.author_id = a.author_id WHERE a.author_name = ?", "Jane Smith")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&b.BookTitle)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		books = append(books, b)
	}

	return books, nil
}

func GetListAllSales(db *sql.DB) (books []entity.Books, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	var b = entity.Books{}

	rows, err := db.QueryContext(ctx, "SELECT b.book_type, SUM(b.price) FROM Order_Details od JOIN Books b ON od.book_id = b.book_id GROUP BY b.book_type")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&b.BookType, &b.Price)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		books = append(books, b)
	}

	return books, nil
}

func GetCustomer(db *sql.DB) (customers []entity.Customers, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	var c = entity.Customers{}

	rows, err := db.QueryContext(ctx, "SELECT c.customer_name, COUNT(o.customer_id) FROM Orders o JOIN Customers c ON o.customer_id = c.customer_id GROUP BY c.customer_name HAVING COUNT(o.customer_id) > 1")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&c.CustomerName, &c.OrderCount)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}

	return customers, nil
}

func GetTopAuthor(db *sql.DB) (authors entity.Authors, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	err = db.QueryRowContext(ctx, "SELECT a.author_name, SUM(b.price) FROM Order_Details od JOIN Books b ON od.book_id = b.book_id JOIN Authors a ON b.author_id = a.author_id GROUP BY a.author_name ORDER BY SUM(b.price) DESC LIMIT 1").Scan(&authors.AuthorName, &authors.TotalPrice)
	if err != nil {
		fmt.Println(err.Error())
		return authors, err
	}

	return authors, nil
}