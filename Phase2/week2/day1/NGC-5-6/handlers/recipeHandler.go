package handlers

import (
	"Phase2/week2/day1/NGC-5-6/config"
	"Phase2/week2/day1/NGC-5-6/models"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetAllRecipes(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := config.Db.Query("SELECT recipe_id, name, description, cookingTime, rating FROM Recipes")
	if err != nil {
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Status:  http.StatusInternalServerError,
			Message: "Failed to get all recipes" + err.Error(),
		})
		return
	}

	var recipes []models.Recipe

	for rows.Next() {
		var rc models.Recipe
		err = rows.Scan(&rc.RecipeID, &rc.Name, &rc.Description, &rc.CookingTime, &rc.Rating)
		if err != nil {
			json.NewEncoder(w).Encode(models.ErrorMessage{
				Status:  http.StatusInternalServerError,
				Message: "Failed to get all recipes" + err.Error(),
			})
			return
		}
		recipes = append(recipes, rc)
	}

	var success_message = models.SuccessMessage{
		Status:  http.StatusOK,
		Message: "Success to get all recipes",
		Datas:   recipes,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(success_message)
}

func GetRecipeById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	id := ps.ByName("id")

	statement, err := config.Db.Prepare("SELECT recipe_id, name, description, cookingTime, rating FROM Recipes WHERE recipe_id = ?")
	if err != nil {
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Status:  http.StatusInternalServerError,
			Message: "Failed to preare query, " + err.Error(),
		})
		return
	}
	defer statement.Close()

	var rc models.Recipe

	err = statement.QueryRow(id).Scan(&rc.RecipeID, &rc.Name, &rc.Description, &rc.CookingTime, &rc.Rating)
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
		Datas:   rc,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(success_msg)
}

func InsertNewRecipe(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	v := ValidateRole(w, r)
	if v != nil {
		json.NewEncoder(w).Encode(v)
		return
	}

	var rc models.Recipe
	err := json.NewDecoder(r.Body).Decode(&rc)
	if err != nil {
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Status:  http.StatusBadRequest,
			Message: "Invalid body input, " + err.Error(),
		})
		return
	}

	statement, err := config.Db.Prepare("INSERT INTO Recipes (name, description, cookingTime,  rating) VALUES (?,?,?,?)")
	if err != nil {
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Status:  http.StatusInternalServerError,
			Message: "Failed to prepare query, " + err.Error(),
		})
		return
	}
	defer statement.Close()

	_, err = statement.Exec(rc.Name, rc.Description, rc.CookingTime, rc.Rating)
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

func UpdateRecipe(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	v := ValidateRole(w, r)
	if v != nil {
		json.NewEncoder(w).Encode(v)
		return
	}

	id := ps.ByName("id")

	var rc models.Recipe
	err := json.NewDecoder(r.Body).Decode(&rc)
	if err != nil {
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Status:  http.StatusBadRequest,
			Message: "Invalid body input, " + err.Error(),
		})
		return
	}

	statement, err := config.Db.Prepare("UPDATE Recipes SET name = ?, description = ?, cookingTime = ?, rating = ? WHERE recipe_id = ?")
	if err != nil {
		json.NewEncoder(w).Encode(models.ErrorMessage{
			Status:  http.StatusInternalServerError,
			Message: "Failed to prepare query, " + err.Error(),
		})
		return
	}
	defer statement.Close()

	_, err = statement.Exec(rc.Name, rc.Description, rc.CookingTime, rc.Rating, id)
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

func DeleteRecipe(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	v := ValidateRole(w, r)
	if v != nil {
		json.NewEncoder(w).Encode(v)
		return
	}

	id := ps.ByName("id")

	statement, err := config.Db.Prepare("DELETE FROM Recipes WHERE recipe_id = ?")
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

func ValidateRole(w http.ResponseWriter, r *http.Request) interface{} {
	c, err := r.Cookie("role")
	if err != nil {
		return models.ErrorMessage{
			Status:  http.StatusUnauthorized,
			Message: "unauthorized: , " + err.Error(),
		}
	}
	if c.Value != "superadmin" {
		return models.ErrorMessage{
			Status:  http.StatusUnauthorized,
			Message: "required superadmin role",
		}
	}

	return nil
}
