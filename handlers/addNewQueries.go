package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"server/models"
	"time"
)

func AddQuery(w http.ResponseWriter, r *http.Request) {
	var newQuery models.Query

	if err := json.NewDecoder(r.Body).Decode(&newQuery); err != nil {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	path, err := filepath.Abs("database.json")
	if err != nil {
		http.Error(w, "Failed to find database file", http.StatusInternalServerError)
		return
	}

	data, err := os.ReadFile(path)
	if err != nil {
		http.Error(w, "Failed to read database", http.StatusInternalServerError)
		return
	}

	var existingQueries []models.Query
	if err := json.Unmarshal(data, &existingQueries); err != nil {
		http.Error(w, "Invalid database format", http.StatusInternalServerError)
		return
	}

	newQuery.ID = len(existingQueries) + 1
	newQuery.SubmittedAt = time.Now().UTC().Format(time.RFC3339)
	if newQuery.Status == "" {
		newQuery.Status = "Pending"
	}

	existingQueries = append(existingQueries, newQuery)

	updatedData, err := json.MarshalIndent(existingQueries, "", "  ")
	if err != nil {
		http.Error(w, "Failed to encode updated data", http.StatusInternalServerError)
		return
	}

	err = os.WriteFile(path, updatedData, 0644)
	if err != nil {
		http.Error(w, "Failed to save query", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Query submitted successfully"})
}
