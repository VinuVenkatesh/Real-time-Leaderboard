package middleware

import (
	"net/http"
	"os"
	"strings"
)

func CORS(next http.Handler) http.Handler {
	allowedOrigins := strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ",")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestOrigin := r.Header.Get("Origin")
		allowed := false
		for _, origin := range allowedOrigins {
			if origin == requestOrigin {
				allowed = true
				w.Header().Set("Access-Control-Allow-Origin", requestOrigin)
				break
			}
		}
		if !allowed {
			http.Error(w, "CORS origin not allowed", http.StatusForbidden)
			return
		}
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
