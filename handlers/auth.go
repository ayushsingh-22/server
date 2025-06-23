package handlers

import (
	"encoding/json"
	"net/http"
	"server/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("your_secret_key") // Use a secure, random key in production

var adminUser = models.Admin{
	Email:    "qwerty@gmail.com",
	Password: "qwety",
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var login models.Admin
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	if login.Email == adminUser.Email && login.Password == adminUser.Password {
		// Create JWT token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email": login.Email,
			"exp":   time.Now().Add(24 * time.Hour).Unix(),
		})
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Could not generate token"})
			return
		}

		// Option 1: Send token in response body
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Login successful",
			"token":   tokenString,
		})

		// Option 2: (Optional) Set token as HttpOnly cookie
		// http.SetCookie(w, &http.Cookie{
		// 	Name:     "token",
		// 	Value:    tokenString,
		// 	Path:     "/",
		// 	HttpOnly: true,
		// 	SameSite: http.SameSiteLaxMode,
		// 	Expires:  time.Now().Add(24 * time.Hour),
		// })
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(map[string]string{"error": "Invalid credentials"})
}

func CheckLoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"authenticated": true}`))
}

// Middleware to protect routes
func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Option 1: Get token from Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}
		tokenString := authHeader[len("Bearer "):]

		// Option 2: (If using cookie) Get token from cookie
		// cookie, err := r.Cookie("token")
		// if err != nil {
		// 	http.Error(w, "Missing token cookie", http.StatusUnauthorized)
		// 	return
		// }
		// tokenString := cookie.Value

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
