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
	var login models.Admin
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if login.Email == adminUser.Email && login.Password == adminUser.Password {
		// Set login cookie; you may generate a token for production use.
		cookie := &http.Cookie{
			Name:     "session",
			Value:    "authenticate",
			Path:     "/",
			HttpOnly: false,
			SameSite: http.SameSiteLaxMode,
		}
		http.SetCookie(w, cookie)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Login successful",
		})
		return
	}

	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// To logout, clear the cookie by setting MaxAge to -1.
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Logout successful",
	})
}
