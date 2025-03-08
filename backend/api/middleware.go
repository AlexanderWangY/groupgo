package api

import (
	"log"
	"net/http"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the access_token from cookies
		cookie, err := r.Cookie("access_token")
		if err != nil {
			// If the cookie is not found, return an error
			http.Error(w, "No access_token cookie.", http.StatusUnauthorized)
			return
		}

		// Optionally, you can validate the token here (e.g., JWT decoding)
		accessToken := cookie.Value
		if accessToken == "" {
			http.Error(w, "Invalid access_token.", http.StatusUnauthorized)
			return
		}

		log.Printf("Token: %s", accessToken)

		// Proceed with the next handler if token is valid
		next.ServeHTTP(w, r)
	})
}
