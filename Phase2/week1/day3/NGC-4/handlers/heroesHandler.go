package handlers

import (
	"Phase2/week1/day3/NGC-4/config"
	"Phase2/week1/day3/NGC-4/models"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetHeroes(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
