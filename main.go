package main

import (
	"log"
	"net/http"
	"os"

	"server/handlers" // ensure this import path matches your folder structure

	"github.com/rs/cors"
)

func main() {

	if os.Getenv("GEMINI_API_KEY") == "" {
		log.Println("Warning: GEMINI_API_KEY environment variable not set. Chatbot functionality will not work properly.")
	}
	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})
	mux.HandleFunc("/api/login", handlers.LoginHandler)
	mux.HandleFunc("/api/logout", handlers.LogoutHandler)
	mux.HandleFunc("/api/getAllQueries", handlers.GetAllQueries)
	mux.HandleFunc("/api/add-query", handlers.AddQuery)
	mux.HandleFunc("/api/updateStatus", handlers.UpdateQueryStatus)
	mux.HandleFunc("/api/check-login", handlers.CheckLoginStatus)
	mux.HandleFunc("/api/analytics", handlers.AnalyticsHandler)
	mux.HandleFunc("/api/chat", handlers.ChatHandler)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://server-saby.onrender.com"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
	})

	// Start server
	log.Println("Server is running on onrender")
	if err := http.ListenAndServe(":8080", c.Handler(mux)); err != nil {
		log.Fatal(err)
	}
}
