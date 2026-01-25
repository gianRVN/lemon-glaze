package handlers

import (
	"encoding/json"
	"net/http"
	"categories/utils"
	"categories/data"
	"categories/types"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.Categories)
}

func GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIDFromPath(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	for _, c := range data.Categories {
		if c.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(c)
			return
		}
	}

	http.Error(w, "Category not found", http.StatusNotFound)
}

func PostCategory(w http.ResponseWriter, r *http.Request) {
	var newCategory types.Category
	if err := json.NewDecoder(r.Body).Decode(&newCategory); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newCategory.ID = len(data.Categories) + 1
	data.Categories = append(data.Categories, newCategory)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCategory)
}

func PutCategory(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIDFromPath(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	var updatedCategory types.Category
	err = json.NewDecoder(r.Body).Decode(&updatedCategory)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	for i := range data.Categories {
		if data.Categories[i].ID == id {
			updatedCategory.ID = id
			data.Categories[i] = updatedCategory

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedCategory)
			return
		}
	}

	http.Error(w, "Category not found", http.StatusNotFound)
}


func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIDFromPath(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	for i, c := range data.Categories {
		if c.ID == id {
			data.Categories = append(
				data.Categories[:i],
				data.Categories[i+1:]...,
			)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "success delete",
			})
			return
		}
	}

	http.Error(w, "Category not found", http.StatusNotFound)
}
