package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("secreto")

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
