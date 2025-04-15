package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
)

type Query struct {
	// Define fields based on the structure of your JSON data
	// Example:
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

func GetAllQueries(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs("database.json")
	if err != nil {
		http.Error(w, "Failed to resolve file path", http.StatusInternalServerError)
		return
	}

	data, err := os.ReadFile(path)
	if err != nil {
		http.Error(w, "Unable to read database file", http.StatusInternalServerError)
		return
	}

	var queries []Query
	if err := json.Unmarshal(data, &queries); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(queries)
}
