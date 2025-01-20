package jwtutil

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var secretKey = []byte("k2OaEi91ZkWj1pBB50dZkxE2LkMvsG5l")

// ExtractUserID extracts the user ID from a JWT token
func ExtractUserID(tokenString string) (int32, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if userIDStr, ok := claims["sub"].(string); ok {
			// Convert userID to int32
			var userID int32
			_, err := fmt.Sscanf(userIDStr, "%d", &userID)
			if err != nil {
				return 0, errors.New("invalid user ID format")
			}
			return userID, nil
		}
	}
	return 0, errors.New("invalid token or missing user ID")
}

// IsTokenValid checks if the JWT token is valid
func IsTokenValid(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})

	// Return false if an error occurs during parsing
	if err != nil {
		return false
	}

	// Token is valid if claims are properly parsed and token is not expired
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check the expiration claim 'exp'
		if exp, ok := claims["exp"].(float64); ok {
			// Check if the token is expired (expiration claim 'exp' is a Unix timestamp)
			if exp < float64(time.Now().Unix()) {
				return false
			}
		}
		return true
	}

	// Return false if token claims are invalid
	return false
}
