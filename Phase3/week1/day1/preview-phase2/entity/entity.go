package entity

type User struct {
	ID       int     `json:"user_id,omitempty" gorm:"column:user_id;primaryKey"`
	Email    string  `json:"email" gorm:"not null;unique"`
	Password string  `json:"password,omitempty" gorm:"not null"`
	Balance  float64 `json:"balance" gorm:"not null;default:0"`
}

type Loan struct {
	ID     int     `json:"loan_id,omitempty" gorm:"column:loan_id;primaryKey"`
	UserID int     `json:"user_id" gorm:"primaryKey"`
	User   User    `json:"user,omitempty"`
	Limit  float64 `json:"limit" gorm:"not null;default:0"`
}
