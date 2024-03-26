package handlers

import (
	"Phase2/week2/day1/preview-week1/config"
	"Phase2/week2/day1/preview-week1/models"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetAllBranches(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := config.Db.Query("SELECT branch_id, name, location FROM Branches")
	if err != nil {
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Status:  http.StatusInternalServerError,
			Message: "Failed to get all branches" + err.Error(),
		})
		return
	}

	var branches []models.Branch

	for rows.Next() {
		var b models.Branch
		err = rows.Scan(&b.BranchID, &b.Name, &b.Location)
		if err != nil {
			json.NewEncoder(w).Encode(models.ErrorMessage{
				Status:  http.StatusInternalServerError,
				Message: "Failed to get all products" + err.Error(),
			})
			return
		}
		branches = append(branches, b)
	}

	var success_message = models.SuccessMessage{
		Status:  http.StatusOK,
		Message: "Success to get all branches",
		Datas:   branches,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(success_message)
}

func GetBranchById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	id := ps.ByName("id")

	statement, err := config.Db.Prepare("SELECT branch_id, name, location FROM Branches WHERE branch_id = ?")
	if err != nil {
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Status:  http.StatusInternalServerError,
			Message: "Failed to preare query, " + err.Error(),
		})
		return
	}
	defer statement.Close()

	var b models.Branch

	err = statement.QueryRow(id).Scan(&b.BranchID, &b.Name, &b.Location)
	if err != nil {
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Status:  http.StatusInternalServerError,
			Message: "Failed to scan query, " + err.Error(),
		})
		return
	}
	defer statement.Close()

	var success_msg = models.SuccessMessage{
		Status:  http.StatusOK,
		Message: "Success get data by id",
		Datas:   b,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(success_msg)
}

func InsertNewBranches(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	var b models.Branch
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Status:  http.StatusBadRequest,
			Message: "Invalid body input, " + err.Error(),
		})
		return
	}

	statement, err := config.Db.Prepare("INSERT INTO Branches (name, location) VALUES (?,?)")
	if err != nil {
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Status:  http.StatusInternalServerError,
			Message: "Failed to prepare query, " + err.Error(),
		})
		return
	}
	defer statement.Close()

	_, err = statement.Exec(b.Name, b.Location)
	if err != nil {
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Status:  http.StatusInternalServerError,
			Message: "Failed to insert new data, " + err.Error(),
		})
		return
	}

	var success_msg = models.SuccessMessage{
		Status:  http.StatusCreated,
		Message: "Success insert new data",
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(success_msg)
}

func UpdateBranch(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	id := ps.ByName("id")

	var b models.Branch
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Status:  http.StatusBadRequest,
			Message: "Invalid body input, " + err.Error(),
		})
		return
	}

	statement, err := config.Db.Prepare("UPDATE Branches SET name = ?, location = ? WHERE branch_id = ?")
	if err != nil {
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Status:  http.StatusInternalServerError,
			Message: "Failed to prepare query, " + err.Error(),
		})
		return
	}
	defer statement.Close()

	_, err = statement.Exec(b.Name, b.Location, id)
	if err != nil {
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Status:  http.StatusInternalServerError,
			Message: "Failed to insert id not found, " + err.Error(),
		})
		return
	}

	var success_msg = models.SuccessMessage{
		Status:  http.StatusOK,
		Message: "Success update data by id",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(success_msg)
}

func DeleteBranch(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	id := ps.ByName("id")

	statement, err := config.Db.Prepare("DELETE FROM Branches WHERE branch_id = ?")
	if err != nil {
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Status:  http.StatusInternalServerError,
			Message: "Failed to prepare query, " + err.Error(),
		})
		return
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Status:  http.StatusInternalServerError,
			Message: "Failed to delete id not found, " + err.Error(),
		})
		return
	}

	var success_msg = models.SuccessMessage{
		Status:  http.StatusOK,
		Message: "Success delete data by id",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(success_msg)
}
