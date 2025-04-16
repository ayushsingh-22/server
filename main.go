package main

import (
	"log"
	"net/http"

	"server/handlers" // ensure this import path matches your folder structure

	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})
	mux.HandleFunc("/api/login", handlers.LoginHandler)
	mux.HandleFunc("/api/getAllQueries", handlers.GetAllQueries)
	mux.HandleFunc("/api/add-query", handlers.AddQuery)
	mux.HandleFunc("/api/updateStatus", handlers.UpdateQueryStatus)
	mux.HandleFunc("/api/check-login", handlers.CheckLoginStatus) // ✅ Check login handler
	mux.HandleFunc("/api/analytics", handlers.AnalyticsHandler)

	// ✅ CORS setup for cookie-based auth from localhost:3000
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // ✅ exact frontend origin
		AllowCredentials: true,                              // ✅ allow cookies
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
	})

	// Start server
	log.Println("🚀 Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", c.Handler(mux)); err != nil {
		log.Fatal(err)
	}
}
