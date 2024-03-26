package handlers

import (
	"Phase2/week1/day4/deploy-webserver/config"
	"Phase2/week1/day4/deploy-webserver/models"
	"encoding/json"
	"net/http"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	// decode body request then assign to models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	statement, err := config.DB.Prepare("INSERT INTO books(title, author, year) VALUES (?, ?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer statement.Close()

	result, err := statement.Exec(book.Title, book.Author, book.Year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get id that generated when insert data
	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return response
	book.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}
