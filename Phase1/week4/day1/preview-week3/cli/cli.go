package cli

import (
	"Phase1/week4/day1/preview-week3/config"
	"Phase1/week4/day1/preview-week3/handler"
	"database/sql"
	"fmt"
	"os"
)

func RunApplication() {
	db, err := config.GetConnection()
	if err != nil {
		fmt.Println(err.Error())
		return 
	}
	defer db.Close()

	command := os.Args[1:]

	switch command[0] {
	case "books":
		ShowBooks(db)
	case "sales":
		ShowSales(db)
	case "customers":
		ShowCustomer(db)
	case "topauthor":
		ShowTopAuthor(db)
	default:
		fmt.Println("Unkown Command")
	}
}

func ShowBooks(db *sql.DB)  {
	books, err := handler.GetListAllBooks(db)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Books by Jane Smith: ")
	for _, v := range books {
		fmt.Println(v.BookTitle)
	}
}

func ShowSales(db *sql.DB)  {
	sales, err := handler.GetListAllSales(db)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Total sales for each book type: ")
	for _, v := range sales {
		fmt.Println(v.BookType, ": ", v.Price)
	}
}

func ShowCustomer(db *sql.DB)  {
	customer, err := handler.GetCustomer(db)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Customer who ordered more than 1 book: ")
	for _, v := range customer {
		fmt.Println(v.CustomerName, ": ", v.OrderCount, "orders")
	}
}

func ShowTopAuthor(db *sql.DB)  {
	authors, err := handler.GetTopAuthor(db)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Top earning author:", authors.AuthorName, "with earnings:", authors.TotalPrice)
}