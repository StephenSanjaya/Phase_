package models

type Product struct {
	ProductID   int     `json:"product_id"`
	StoreID     int     `json:"store_id" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required,min=10"`
	ImageUrl    string  `json:"image_url" binding:"required"`
	Price       float64 `json:"price" binding:"required,gte=1000"`
}

type Login struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=5"`
}

type Store struct {
	StoreID   int    `json:"store_id"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=5"`
	StoreName string `json:"store_name" binding:"required,min=6,max=15"`
	StoreType string `json:"store_type" binding:"required"`
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
