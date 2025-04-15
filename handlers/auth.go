package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"server/models"
	"time"

	"github.com/google/uuid"
	"github.com/rs/cors"
)

var adminUser = models.Admin{
	Email:    "qwerty@gmail.com",
	Password: "qwerty",
}

// LoginHandler handles the login logic.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
	})
	handler := c.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var login models.Admin
		log.Println("LoginHandler called")
		err := json.NewDecoder(r.Body).Decode(&login)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Invalid request payload",
			})
			return
		}

		if login.Email == adminUser.Email && login.Password == adminUser.Password {
			// Generate a secure token using uuid.
			token := uuid.New().String()
			cookie := &http.Cookie{
				Name:     "auth",
				Value:    token, // Token generated from uuid.New().String()
				Path:     "/",
				HttpOnly: true,
				Expires:  time.Now().Add(1 * time.Hour),
			}
			http.SetCookie(w, cookie)

			// Printing token value and the cookie storage path
			log.Println("Token value:", token)
			log.Println("Cookie stored at path:", cookie.Path)

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Login successful",
			})
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Invalid credentials",
			})
		}
	}))
	handler.ServeHTTP(w, r)
}

// LogoutHandler handles the logout logic.
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Clear the auth cookie by setting its expiry to the past.
	cookie := &http.Cookie{
		Name:     "auth",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	}
	http.SetCookie(w, cookie)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Logout successful",
	})
}
