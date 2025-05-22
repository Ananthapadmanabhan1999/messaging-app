package auth

import (
	"net/http"
	"strings"
)

func AuthMiddleware(service *AuthService, next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if !strings.HasPrefix(authHeader, "Bearer ") {
            http.Error(w, "missing token", http.StatusUnauthorized)
            return
        }

        token := strings.TrimPrefix(authHeader, "Bearer ")
        _, ok := service.ValidateToken(token)
        if !ok {
            http.Error(w, "invalid token", http.StatusUnauthorized)
            return
        }

        next(w, r)
    }
}