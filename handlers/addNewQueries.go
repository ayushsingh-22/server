package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"server/models"

	)

	
func AddQuery(w http.ResponseWriter, r *http.Request) {
	// Decode new query from request body
	var newQuery models.Query
	if err := json.NewDecoder(r.Body).Decode(&newQuery); err != nil {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	// Load existing data
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

	// Generate new ID and set the submitted time.
	newQuery.ID = len(existingQueries) + 1
	newQuery.SubmittedAt = time.Now().UTC().Format(time.RFC3339)

	// Append new query
	existingQueries = append(existingQueries, newQuery)

	// Save back to file
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

	// Return success
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Query submitted successfully"})
}
