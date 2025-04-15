package main

import (
	"log"
	"net/http"

	"server/handlers"

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

	// Allow all origins with CORS
	c := cors.AllowAll()

	log.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", c.Handler(mux)); err != nil {
		log.Fatal(err)
	}
}
