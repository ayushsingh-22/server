package handlers

import (
	"encoding/json"
	"net/http"
	"server/models"
)

var adminUser = models.Admin{
	Email:    "qwerty@gmail.com",
	Password: "qwety",
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type for consistent JSON responses
	w.Header().Set("Content-Type", "application/json")

	var login models.Admin
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	if login.Email == adminUser.Email && login.Password == adminUser.Password {
		// Set login cookie; you may generate a token for production use.
		cookie := &http.Cookie{
			Name:     "session",
			Value:    "authenticated",
			Path:     "/",
			HttpOnly: false, // Set to true in production for security
			SameSite: http.SameSiteLaxMode,
			MaxAge:   3600 * 24, // 24 hours
		}
		http.SetCookie(w, cookie)

		json.NewEncoder(w).Encode(map[string]string{
			"message": "Login successful",
			"status":  "authenticated",
		})
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(map[string]string{"error": "Invalid credentials"})
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type
	w.Header().Set("Content-Type", "application/json")

	// Clear the session cookie
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Logout successful",
		"status":  "logged_out",
	})
}

func CheckLoginStatus(w http.ResponseWriter, r *http.Request) {
	// Set content type
	w.Header().Set("Content-Type", "application/json")

	// Check for session cookie
	cookie, err := r.Cookie("session")

	if err != nil || cookie.Value != "authenticated" {
		// No valid session found
		w.WriteHeader(http.StatusOK) // Still 200 OK
		json.NewEncoder(w).Encode(map[string]bool{
			"authenticated": false,
		})
		return
	}

	// Valid session found
	json.NewEncoder(w).Encode(map[string]bool{
		"authenticated": true,
	})
}
