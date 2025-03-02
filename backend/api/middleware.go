package api

import "net/http"

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if authorization == "" {
			http.Error(w, "No authorization header.", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
