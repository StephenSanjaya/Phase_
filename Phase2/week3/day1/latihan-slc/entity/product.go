package entity

type Product struct {
	ID             uint    `gorm:"primaryKey" json:"id,omitempty"`
	Name           string  `gorm:"not null" json:"name" binding:"required"`
	Description    string  `gorm:"not null" json:"description" binding:"required"`
	Price          float32 `gorm:"not null" json:"price" binding:"required"`
	InventoryLevel uint    `gorm:"not null" json:"inventory_level" binding:"required"`
}
