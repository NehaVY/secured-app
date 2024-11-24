package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/echo", authMiddleware(echoHandler))
	http.ListenAndServe(":8080", nil)
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Query().Get("input")

	// Secure: Input sanitization to prevent injection
	if len(input) > 50 || input == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	response := map[string]string{"message": fmt.Sprintf("Echo: %s", input)}
	json.NewEncoder(w).Encode(response) // Secure: JSON response
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Secure: Basic token-based authentication
		token := r.Header.Get("Authorization")
		if token != "Bearer securetoken123" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

// Secure: Environment variable fetched securely
func getEnvVariable() string {
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		panic("SECRET_KEY not set")
	}
	return secretKey
}
