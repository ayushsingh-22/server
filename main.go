package main

import (
	"log"
	"net/http"

	"server/handlers" // ensure this import is correct

	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	mux.HandleFunc("/api/login", handlers.LoginHandler)
	mux.HandleFunc("/api/getAllQueries", handlers.GetAllQueries)
	mux.HandleFunc("/api/logout", handlers.LogoutHandler)
	mux.HandleFunc("/api/add-query", handlers.AddQuery)
	mux.HandleFunc("/api/updateStatus", handlers.UpdateQueryStatus) // ✅ ADD THIS

	c := cors.AllowAll()

	log.Println("🚀 Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", c.Handler(mux)); err != nil {
		log.Fatal(err)
	}
}
