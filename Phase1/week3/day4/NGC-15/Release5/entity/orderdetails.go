package entity

type OrderDetails struct {
	OrderID, TableNumber int
	Name string
	TotalPrice float64
}