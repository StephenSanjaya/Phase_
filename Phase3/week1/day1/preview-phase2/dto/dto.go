package dto

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Loan struct {
	Limit float64 `json:"limit"`
}

type WithdrawBalance struct {
	Balance float64 `json:"balance"`
}
