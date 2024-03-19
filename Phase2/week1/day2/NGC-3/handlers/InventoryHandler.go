package handlers

import (
	"Phase2/week1/day2/NGC-3/config"
	"Phase2/week1/day2/NGC-3/models"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetAllInventories(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rows, err := config.DB.Query("SELECT id, name, itemcode, stock, description, status FROM inventories")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var inventories []models.Inventory
	for rows.Next() {
		var inven models.Inventory

		err = rows.Scan(&inven.ID, &inven.Name, &inven.ItemCode, &inven.Stock, &inven.Description, &inven.Status)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		inventories = append(inventories, inven)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inventories)
}

func GetInventoryByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	var inven models.Inventory

	err := config.DB.QueryRow("SELECT id, name, itemcode, stock, description, status FROM inventories WHERE id = ?", id).Scan(&inven.ID, &inven.Name, &inven.ItemCode, &inven.Stock, &inven.Description, &inven.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inven)
}

func CreateItem(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var inven models.Inventory

	err := json.NewDecoder(r.Body).Decode(&inven)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	statement, err := config.DB.Prepare("INSERT INTO inventories(name, itemcode, stock, description, status) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer statement.Close()

	result, err := statement.Exec(inven.Name, inven.ItemCode, inven.Stock, inven.Description, inven.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	inven.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inven)
}

func UpdateInventoryById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var inven models.Inventory

	err := json.NewDecoder(r.Body).Decode(&inven)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := ps.ByName("id")
	_, err = config.DB.Exec("UPDATE inventories SET name = ?, itemcode = ?, stock = ?, description = ?, status = ? WHERE id = ?", inven.Name, inven.ItemCode, inven.Stock, inven.Description, inven.Status, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Book updated successfully")
}

func DeleteInventoryById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	_, err := config.DB.Exec("DELETE FROM inventories WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Book deleted successfully")
}
