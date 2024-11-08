package middleware

import (
	"Hotelsystem/services"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		_, err := services.ValidateToken(tokenStr)
		if err != nil {
			http.Error(w, "token invalido", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
