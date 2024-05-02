package models

type Item struct {
	ID    int     `json:"id,omitempty"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
