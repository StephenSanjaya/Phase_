package entity

type Order struct {
	ID                  uint   `gorm:"primaryKey" json:"id,omitempty"`
	CreatedAt           string `gorm:"type:date;column:order_date;not null" json:"order_date" binding:"required"`
	CustomerID          uint   `gorm:"not null" json:"customer_id" binding:"required"`
	PaymentMethod       string `gorm:"not null" json:"payment_method" binding:"required"`
	PaymentConfirmation string `gorm:"not null" json:"payment_confirmation" binding:"required"`
	Payments            []Payment
	Products            []Product `gorm:"many2many:order_items"`
}

type OrderItem struct {
	OrderID   uint `gorm:"primaryKey"`
	ProductID uint `gorm:"primaryKey"`
	Quantity  uint `gorm:"not null"`
}
