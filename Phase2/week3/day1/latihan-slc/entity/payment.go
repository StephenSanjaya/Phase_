package entity

type Payment struct {
	ID        uint    `gorm:"primaryKey" json:"id,omitempty"`
	OrderID   uint    `gorm:"not null" json:"order_id" binding:"required"`
	Amount    float64 `gorm:"not null" json:"amount" binding:"required"`
	CreatedAt string  `gorm:"type:date;column:payment_date;not null" json:"payment_date" binding:"required"`
}
