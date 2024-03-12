package entity

type Books struct {
	BookTitle, BookType string
	Price float64
}

type Customers struct {
	CustomerName, CustomerEmail string
	OrderCount int
}

type Authors struct {
	AuthorName, AuthoEmail string
	TotalPrice float64
}
