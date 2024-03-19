package handlers

import (
	"Phase2/week1/day2/go-routing/config"
	"Phase2/week1/day2/go-routing/models"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var book models.Book

	// decode body request then assign to models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := ps.ByName("id")
	_, err = config.DB.Exec("UPDATE books SET title = ?, author = ?, year = ? WHERE id = ?", book.Title, book.Author, book.Year, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Book updated successfully")
}
