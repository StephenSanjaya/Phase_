package entity

type Payment struct {
	OrderID int
	PaymentDate, PaymentMethod string
	TotalAmount float64
}