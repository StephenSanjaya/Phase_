package handlers

import (
	"Phase2/week1/day4/deploy-ngc4/config"
	"Phase2/week1/day4/deploy-ngc4/models"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
}

func GetAllCriminalReports(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	enableCors(&w)

	rows, err := config.DB.Query("SELECT criminal_report_id, hero_id, villain_id, description, date FROM CriminalReports")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Message: "Failed to query",
			Status:  http.StatusInternalServerError,
		})
		return
	}
	defer rows.Close()

	var criminal_reports []models.CriminalReports
	w.Header().Set("Content-Type", "application/json")

	for rows.Next() {
		var r models.CriminalReports

		err = rows.Scan(&r.ID, &r.HeroID, &r.VillainID, &r.Description, &r.Date)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.ErrorMessage{
				Message: "Failed to query",
				Status:  http.StatusInternalServerError,
			})
			return
		}

		criminal_reports = append(criminal_reports, r)
	}

	var success_msg = models.SuccessMessage{
		Status:  http.StatusOK,
		Message: "Show Data Criminal Reports",
		Datas:   criminal_reports,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(success_msg)
}

func GetCriminalReportById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	enableCors(&w)

	id := ps.ByName("id")
	var cr models.CriminalReports
	w.Header().Set("Content-Type", "application/json")

	statement, err := config.DB.Prepare("SELECT criminal_report_id, hero_id, villain_id, description, date FROM CriminalReports WHERE criminal_report_id = ?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Message: "Failed to query",
			Status:  http.StatusInternalServerError,
		})
		return
	}
	defer statement.Close()

	err = statement.QueryRow(id).Scan(&cr.ID, &cr.HeroID, &cr.VillainID, &cr.Description, &cr.Date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Message: "Failed to query",
			Status:  http.StatusInternalServerError,
		})
		return
	}

	var success_msg = models.SuccessMessage{
		Message: "Show Data Criminal Reports By Id",
		Status:  http.StatusOK,
		Data:    &cr,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(success_msg)
}

func CreateReport(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	enableCors(&w)

	var cr models.CriminalReports

	err := json.NewDecoder(r.Body).Decode(&cr)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Message: "Invalid body input",
			Status:  http.StatusBadRequest,
		})
		return
	}

	statement, err := config.DB.Prepare("INSERT INTO CriminalReports(hero_id, villain_id, description, date) VALUES (?, ?, ?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Message: "Failed to query",
			Status:  http.StatusInternalServerError,
		})
		return
	}
	defer statement.Close()

	result, err := statement.Exec(cr.HeroID, cr.VillainID, cr.Description, cr.Date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Message: "Failed to query",
			Status:  http.StatusInternalServerError,
		})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Message: "Failed to query",
			Status:  http.StatusInternalServerError,
		})
		return
	}

	cr.ID = id
	var success_msg = models.SuccessMessage{
		Message: "success create new report",
		Status:  http.StatusCreated,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(success_msg)
}

func UpdateReportById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	enableCors(&w)

	var cr models.CriminalReports

	err := json.NewDecoder(r.Body).Decode(&cr)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Message: "Invalid body input",
			Status:  http.StatusBadRequest,
		})
		return
	}

	id := ps.ByName("id")
	statement, err := config.DB.Prepare("UPDATE CriminalReports SET hero_id = ?, villain_id = ?, description = ?, date = ? WHERE criminal_report_id = ?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Message: "Failed to query",
			Status:  http.StatusInternalServerError,
		})
		return
	}
	defer statement.Close()

	_, err = statement.Exec(cr.HeroID, cr.VillainID, cr.Description, cr.Date, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Message: "Failed to query",
			Status:  http.StatusInternalServerError,
		})
		return
	}

	var success_msg = models.SuccessMessage{
		Message: "success update report",
		Status:  http.StatusOK,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(success_msg)
}

func DeleteReportById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	enableCors(&w)

	w.Header().Set("Content-Type", "application/json")
	id := ps.ByName("id")

	statement, err := config.DB.Prepare("DELETE FROM CriminalReports WHERE criminal_report_id = ?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Message: "Failed to query delete",
			Status:  http.StatusInternalServerError,
		})
		return
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Message: "Failed to query delete",
			Status:  http.StatusInternalServerError,
		})
		return
	}

	var success_msg = models.SuccessMessage{
		Message: "success delete report",
		Status:  http.StatusOK,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(success_msg)
}
