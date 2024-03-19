package handlers

import (
	"Phase2/week1/day2/go-web-server/config"
	"Phase2/week1/day2/go-web-server/models"
	"encoding/json"
	"net/http"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id, title, author, year FROM books")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book

		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		books = append(books, book)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
