package handlers

import (
	"Phase2/week1/day2/go-routing/config"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	_, err := config.DB.Exec("DELETE FROM books WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Book deleted successfully")
}
