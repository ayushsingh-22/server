package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type Query struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Service         string `json:"service"`
	Guards          int    `json:"guards"`
	DurationValue   int    `json:"durationValue"`
	DurationUnit    string `json:"durationUnit"`
	CameraRequired  bool   `json:"cameraRequired"`
	VehicleRequired bool   `json:"vehicleRequired"`
	SpecialRequest  string `json:"specialRequest"`
	SubmittedAt     string `json:"submitted_at"`
}

func AddBooking(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
		return
	}

	var newQuery Query
	if err := json.NewDecoder(r.Body).Decode(&newQuery); err != nil {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	// Read existing queries from file
	path, err := filepath.Abs("database.json")
	if err != nil {
		http.Error(w, "Failed to find database file", http.StatusInternalServerError)
		return
	}

	var queries []Query
	if data, err := os.ReadFile(path); err == nil {
		_ = json.Unmarshal(data, &queries)
	}

	// Assign ID and timestamp
	newQuery.ID = len(queries) + 1
	newQuery.SubmittedAt = time.Now().UTC().Format(time.RFC3339)

	// Append the new query and save the updated list
	queries = append(queries, newQuery)
	updatedData, err := json.MarshalIndent(queries, "", "  ")
	if err != nil {
		http.Error(w, "Failed to encode updated data", http.StatusInternalServerError)
		return
	}

	if err := os.WriteFile("database.json", updatedData, 0644); err != nil {
		http.Error(w, "Failed to write updated data", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}
