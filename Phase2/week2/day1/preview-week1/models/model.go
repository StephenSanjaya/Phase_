package models

type Branch struct {
	BranchID int    `json:"branch_id"`
	Name     string `json:"name"`
	Location string `json:"location"`
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
