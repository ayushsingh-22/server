package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
)

func GetAllQueries(w http.ResponseWriter, r *http.Request) {
	// Get absolute path to database.json (assuming it's in server root)
	path, err := filepath.Abs("database.json")
	if err != nil {
		http.Error(w, "Failed to resolve file path", http.StatusInternalServerError)
		return
	}

	// Read the file
	data, err := os.ReadFile(path)
	if err != nil {
		http.Error(w, "Unable to read database file", http.StatusInternalServerError)
		return
	}

	// Unmarshal JSON to verify valid structure (optional)
	var queries []map[string]any
	if err := json.Unmarshal(data, &queries); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusInternalServerError)
		return
	}

	// Marshal the queries variable to return it as response.
	response, err := json.MarshalIndent(queries, "", "  ")
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
