package handlers

import (
	"Phase2/week1/day2/NGC-2/config"
	"Phase2/week1/day2/NGC-2/models"
	"encoding/json"
	"net/http"
)

func GetVillain(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id, name, universe, imageUrl FROM villain")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var villains []models.Villain
	for rows.Next() {
		var villain models.Villain

		err = rows.Scan(&villain.ID, &villain.Name, &villain.Universe, &villain.ImageURL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		villains = append(villains, villain)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(villains)
}
