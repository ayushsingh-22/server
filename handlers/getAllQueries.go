package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"

	"server/models"
)

func GetAllQueries(w http.ResponseWriter, r *http.Request) {
	// Get the absolute path of the database file
	path, err := filepath.Abs("database.json")
	if err != nil {
		http.Error(w, "Failed to resolve file path", http.StatusInternalServerError)
		return
	}

	log.Println("Reading database from:", path)

	// Read the file data
	data, err := os.ReadFile(path)
	if err != nil {
		http.Error(w, "Unable to read database file", http.StatusInternalServerError)
		return
	}

	// Unmarshal the data into a slice of models.Query
	var queries []models.Query
	if err := json.Unmarshal(data, &queries); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusInternalServerError)
		log.Println("JSON parsing error:", err)
		return
	}

	log.Printf("Found %d queries before sorting", len(queries))

	// Debug: Log the IDs before sorting
	var idsBeforeSorting []int
	for _, q := range queries {
		idsBeforeSorting = append(idsBeforeSorting, q.ID)
	}
	log.Printf("IDs before sorting: %v", idsBeforeSorting)

	// Sort the queries by ID in DESCENDING order
	sort.SliceStable(queries, func(i, j int) bool {
		return queries[i].ID > queries[j].ID
	})

	// Debug: Log the IDs after sorting
	var idsAfterSorting []int
	for _, q := range queries {
		idsAfterSorting = append(idsAfterSorting, q.ID)
	}
	log.Printf("IDs after sorting: %v", idsAfterSorting)

	// Don't reassign IDs - keep original descending order

	// Set header and response status, then encode and send the sorted queries as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(queries)
}
