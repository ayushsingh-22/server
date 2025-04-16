package handlers

import (
	"encoding/json"
	"net/http"
)

func CheckLoginStatus(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil || cookie.Value == "" {
		http.Error(w, "Not logged in", http.StatusUnauthorized)
		return
	}

	// Optionally: validate session_token if you're storing sessions
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Logged in",
	})
}
