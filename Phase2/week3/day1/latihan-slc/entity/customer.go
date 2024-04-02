package entity

type Customer struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	FirstName   string `gorm:"not null" json:"first_name" binding:"required"`
	LastName    string `gorm:"not null" json:"last_name" binding:"required"`
	Email       string `gorm:"not null;unique" json:"email" binding:"required,email"`
	Password    string `gorm:"not null" json:"password,omitempty" binding:"required"`
	PhoneNumber string `gorm:"not null" json:"phone_number" binding:"required"`
	Address     string `gorm:"not null" json:"address" binding:"required"`
	Orders      []Order `json:"orders,omitempty"`
}
