package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"server/models" // Import the models package, which contains the Query struct defined in admin.go
)

func GetAllQueries(w http.ResponseWriter, r *http.Request) {
	// Get the absolute path of the database file.
	path, err := filepath.Abs("database.json")
	if err != nil {
		http.Error(w, "Failed to resolve file path", http.StatusInternalServerError)
		return
	}

	// Debug: log the resolved path
	log.Println("Resolved database path:", path)

	// Read the file data.
	data, err := os.ReadFile(path)
	if err != nil {
		http.Error(w, "Unable to read database file", http.StatusInternalServerError)
		return
	}

	// Unmarshal the data into a slice of models.Query.
	var queries []models.Query
	if err := json.Unmarshal(data, &queries); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusInternalServerError)
		return
	}

	// Set header and response status, then encode and send the queries as JSON.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(queries)
}
