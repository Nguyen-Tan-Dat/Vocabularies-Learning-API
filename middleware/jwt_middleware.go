package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Nguyen-Tan-Dat/Vocabularies-Learning-API/jwtutil"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Lấy header Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if !jwtutil.IsTokenValid(tokenString) {
			http.Error(w, "Token is valid", http.StatusUnauthorized)
			return
		}

		userID, err := jwtutil.ExtractUserID(tokenString)
		if err != nil {
			http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}
		if userID == 0 {
			http.Error(w, "Invalid userID in token", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "userID", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
