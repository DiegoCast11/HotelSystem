package middleware

import (
	"Hotelsystem/pkg/contextkey"
	"Hotelsystem/services"
	"context"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		token, err := services.ValidateToken(tokenStr)
		if err != nil {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		// Obtener los datos del token (userID, phone)
		claims := token.Claims.(jwt.MapClaims)
		userID := int(claims["user_id"].(float64))
		email := claims["email"].(string)

		// Guardar los valores en el contexto usando las claves personalizadas
		ctx := context.WithValue(r.Context(), contextkey.UserIDKey, userID)
		ctx = context.WithValue(ctx, contextkey.EmailKey, email)

		// Continuar con la siguiente función handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
