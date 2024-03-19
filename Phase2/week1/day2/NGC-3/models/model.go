package models

type Inventory struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	ItemCode    string `json:"itemcode"`
	Stock       int    `json:"stock"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
