package utils

import (
	"errors"
)

// Struktur data Item
type Item struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var items []Item
var lastID int

// CreateItem menambahkan item baru ke slice items.
func CreateItem(item Item) (int, error) {
	lastID++
	item.ID = lastID
	items = append(items, item)
	return item.ID, nil
}

// GetItemByID mengembalikan item berdasarkan ID.
func GetItemByID(id int) (*Item, error) {
	for _, item := range items {
		if item.ID == id {
			return &item, nil
		}
	}
	return nil, errors.New("Item not found")
}

// UpdateItem mengupdate item berdasarkan ID.
func UpdateItem(id int, updatedItem Item) error {
	for i, item := range items {
		if item.ID == id {
			updatedItem.ID = id
			items[i] = updatedItem
			return nil
		}
	}
	return errors.New("Item not found")
}

// DeleteItem menghapus item berdasarkan ID.
func DeleteItem(id int) error {
	for i, item := range items {
		if item.ID == id {
			// Menghapus item dari slice
			items = append(items[:i], items[i+1:]...)
			return nil
		}
	}
	return errors.New("Item not found")
}

// GetAllItems mengembalikan semua item.
func GetAllItems() []Item {
	return items
}
