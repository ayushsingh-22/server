package handlers

import (
	"encoding/json"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"time"

	"server/models"
)

func AnalyticsHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type early to ensure JSON response
	w.Header().Set("Content-Type", "application/json")

	// Get absolute path to database file
	path, err := filepath.Abs("database.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to resolve database path"})
		return
	}

	file, err := os.Open(path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to read database file"})
		return
	}
	defer file.Close()

	var queries []models.Query
	if err := json.NewDecoder(file).Decode(&queries); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to parse JSON data"})
		return
	}

	// Calculate revenue by service and by month
	serviceRevenueMap := make(map[string]float64)
	monthlyRevenueMap := make(map[string]float64)
	layout := "2006-01-02T15:04:05Z" // RFC3339 format without microseconds

	for _, q := range queries {
		// Skip entries with zero cost
		if q.Cost <= 0 {
			continue
		}

		serviceRevenueMap[q.Service] += q.Cost

		// Safely parse the date
		t, err := time.Parse(layout, q.SubmittedAt)
		if err != nil {
			// Try alternate format with timezone name
			t, err = time.Parse("2006-01-02T15:04:05-07:00", q.SubmittedAt)
			if err != nil {
				// Skip this entry if date can't be parsed
				continue
			}
		}

		monthYear := t.Format("Jan 2006")
		monthlyRevenueMap[monthYear] += q.Cost
	}

	// Top 3 Services
	var serviceList []models.TopService
	for k, v := range serviceRevenueMap {
		serviceList = append(serviceList, models.TopService{Service: k, Revenue: v})
	}
	sort.Slice(serviceList, func(i, j int) bool {
		return serviceList[i].Revenue > serviceList[j].Revenue
	})
	topServices := serviceList
	if len(topServices) > 3 {
		topServices = topServices[:3]
	}

	// Pie Chart Data
	var pieChartData []models.ServiceRevenue
	for k, v := range serviceRevenueMap {
		pieChartData = append(pieChartData, models.ServiceRevenue{Name: k, Value: v})
	}

	// Monthly Revenue with Growth
	var monthlyList []models.MonthlyRevenue
	for k, v := range monthlyRevenueMap {
		monthlyList = append(monthlyList, models.MonthlyRevenue{Month: k, Revenue: v})
	}

	// Sort monthly data chronologically
	sort.Slice(monthlyList, func(i, j int) bool {
		ti, err := time.Parse("Jan 2006", monthlyList[i].Month)
		if err != nil {
			return false
		}
		tj, err := time.Parse("Jan 2006", monthlyList[j].Month)
		if err != nil {
			return true
		}
		return ti.Before(tj)
	})

	// Calculate month-over-month growth
	for i := range monthlyList {
		if i > 0 && monthlyList[i-1].Revenue > 0 {
			prev := monthlyList[i-1].Revenue
			curr := monthlyList[i].Revenue
			growth := ((curr - prev) / prev) * 100
			monthlyList[i].Growth = math.Round(growth*100) / 100
		} else {
			monthlyList[i].Growth = 0
		}
	}

	// Final JSON response
	resp := map[string]any{
		"topServices":    topServices,
		"pieChartData":   pieChartData,
		"monthlyRevenue": monthlyList,
	}

	// Encode and send the response
	json.NewEncoder(w).Encode(resp)
}
