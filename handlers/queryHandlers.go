package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

// Query represents a service query record.
type Query struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	Message         string `json:"message"`
	Service         string `json:"service"`
	NumGuards       int    `json:"numGuards"`
	DurationType    string `json:"durationType"`
	DurationValue   int    `json:"durationValue"`
	CameraRequired  bool   `json:"cameraRequired"`
	VehicleRequired bool   `json:"vehicleRequired"`
	SubmittedAt     string `json:"submittedAt"`
	Status          string `json:"status"`
}

type UpdateRequest struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

func UpdateQueryStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // ✅ always set this early

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Only POST method allowed"})
		return
	}

	var req UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	data, err := ioutil.ReadFile("database.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Unable to read database"})
		return
	}

	var queries []Query
	if err := json.Unmarshal(data, &queries); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to parse database"})
		return
	}

	updated := false
	for i := range queries {
		if queries[i].ID == req.ID {
			queries[i].Status = req.Status
			updated = true
			break
		}
	}

	if !updated {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Query not found"})
		return
	}

	updatedData, err := json.MarshalIndent(queries, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to encode updated data"})
		return
	}

	if err := os.WriteFile("database.json", updatedData, 0644); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to write updated data"})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}
