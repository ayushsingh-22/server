package main

import (
	"log"
	"net/http"
	"os"

	"server/handlers" // ensure this import path matches your folder structure

	"github.com/rs/cors"
)

func main() {
	// Set up required environment variables for Gemini API
	if os.Getenv("GEMINI_API_KEY") == "" {
		log.Println("Warning: GEMINI_API_KEY environment variable not set. Chatbot functionality will not work properly.")
	}

	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})
	mux.HandleFunc("/api/login", handlers.LoginHandler)
	mux.HandleFunc("/api/getAllQueries", handlers.GetAllQueries)
	mux.HandleFunc("/api/add-query", handlers.AddQuery)
	mux.HandleFunc("/api/updateStatus", handlers.UpdateQueryStatus)
	mux.HandleFunc("/api/check-login", handlers.CheckLoginStatus)
	mux.HandleFunc("/api/analytics", handlers.AnalyticsHandler)
	
	// Add the new chat endpoint for Gemini integration
	mux.HandleFunc("/api/chat", handlers.ChatHandler)

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