package handlers

import (
	"Phase2/week1/day2/NGC-2/config"
	"Phase2/week1/day2/NGC-2/models"
	"encoding/json"
	"net/http"
)

func GetHeroes(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id, name, universe, skill, imageUrl FROM heroes")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var heroes []models.Heroes
	for rows.Next() {
		var hero models.Heroes

		err = rows.Scan(&hero.ID, &hero.Name, &hero.Universe, &hero.Skill, &hero.ImageURL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		heroes = append(heroes, hero)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(heroes)
}
