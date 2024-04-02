package dto

type Customer struct {
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password,omitempty" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Address     string `json:"address" binding:"required"`
}

type Login struct {
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
}
