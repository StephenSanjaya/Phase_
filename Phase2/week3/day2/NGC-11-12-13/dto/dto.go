package dto

type User struct {
	Username      string  `json:"username"`
	Password      string  `json:"password"`
	DepositAmount float64 `json:"deposit_amount"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Product struct {
	ProductID int     `json:"product_id"`
	Name      string  `json:"name"`
	Stock     int     `json:"stock"`
	Price     float64 `json:"price"`
}

type Transaction struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type Store struct {
	Name      string  `json:"name"`
	Address   string  `json:"address"`
	Longitude string  `json:"longitude,omitempty"`
	Latitude  string  `json:"latitude,omitempty"`
	Rating    float64 `json:"rating,omitempty"`
}
