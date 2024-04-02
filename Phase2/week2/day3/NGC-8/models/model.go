package models

type Product struct {
	ProductID   int     `json:"product_id" gorm:"primaryKey; autoIncrement;not null"`
	StoreID     int     `json:"store_id" binding:"required" gorm:"not null"`
	Name        string  `json:"name" binding:"required" gorm:"not null"`
	Description string  `json:"description" binding:"required,min=10" gorm:"not null"`
	ImageUrl    string  `json:"image_url" binding:"required" gorm:"not null"`
	Price       float64 `json:"price" binding:"required,gte=1000" gorm:"not null"`
}

type Login struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=5"`
}

type Store struct {
	StoreID   int       `json:"store_id" gorm:"primaryKey;autoIncrement;not null"`
	Email     string    `json:"email" binding:"required,email" gorm:"not null"`
	Password  string    `json:"password" binding:"required,min=5" gorm:"not null"`
	StoreName string    `json:"store_name" binding:"required,min=6,max=15" gorm:"not null"`
	StoreType string    `json:"store_type" binding:"required" gorm:"not null"`
	Products  []Product `json:"products"`
}

type ErrorContract struct {
	ErrorCode    int    `json:"code"`
	ErrorMessage string `json:"message"`
	ErrorDetail  string `json:"details"`
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
