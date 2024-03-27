package models

type Product struct {
	ProductID   int     `json:"product_id"`
	StoreID     int     `json:"store_id"`
	Name        string  `json:"name" required:"true"`
	Description string  `json:"description" required:"true" minLen:"8"`
	ImageUrl    string  `json:"image_url" required:"true"`
	Price       float64 `json:"price" required:"true" min:"1000"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Store struct {
	StoreID   int    `required:"true" json:"store_id"`
	Email     string `required:"true" json:"email" regex:"true"`
	Password  string `required:"true" json:"password" minLen:"8"`
	StoreName string `required:"true" json:"store_name" minLen:"6" maxLen:"15"`
	StoreType string `required:"true" json:"store_type"`
}

type SuccessMessage struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Datas   interface{} `json:"datas,omitempty"`
}

type ErrorMessage struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
