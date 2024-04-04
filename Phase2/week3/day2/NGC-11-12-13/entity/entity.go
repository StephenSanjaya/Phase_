package entity

type User struct {
	ID            int       `json:"user_id" gorm:"primaryKey;column:user_id"`
	Username      string    `json:"username" gorm:"type:string;size:255;not null;"`
	Password      string    `json:"password" gorm:"type:string;size:255;not null;"`
	DepositAmount float64   `json:"deposit_amount" gorm:"not null;"`
	Products      []Product `json:"products" gorm:"many2many:transactions"`
}

type Product struct {
	ID    int     `json:"product_id" gorm:"primaryKey;column:product_id"`
	Name  string  `json:"name" gorm:"size:255;not null"`
	Stock int     `json:"stock" gorm:"not null"`
	Price float64 `json:"price" gorm:"not null"`
}

type Transaction struct {
	ID     int `json:"transaction_id" gorm:"primaryKey;column:transaction_id"`
	UserID int `json:"user_id" gorm:"primaryKey;not null;"`
	// StoreID     int     `json:"store_id" gorm:"primaryKey;not null;"`
	ProductID   int     `json:"product_id" gorm:"primaryKey;not null;"`
	Quantity    int     `json:"quantity" gorm:"not null;"`
	TotalAmount float64 `json:"total_amount" gorm:"not null;"`
}

// type Store struct {
// 	ID           int     `json:"transaction_id" gorm:"primaryKey;column:store_id"`
// 	Name         string  `json:"name" gorm:"not null;"`
// 	Address      string  `json:"address" gorm:"not null;"`
// 	Longitude    string  `json:"longitude" gorm:"not null;"`
// 	Latitude     string  `json:"latitude" gorm:"not null;"`
// 	Rating       float64 `json:"rating" gorm:"not null;"`
// 	Transactions []Transaction
// }
