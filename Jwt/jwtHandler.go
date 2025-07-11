package Jwt

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func getSecretKey(keyname string) string {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	key := os.Getenv(keyname)
	return key
}

var jwtKey = []byte(getSecretKey("ACCESS_SECREET"))

func GenerateJWT(id string, exp time.Duration) (string, error) {
	// Define expiration time
	expirationTime := time.Now().Add(exp * time.Hour)

	// Create the claims
	data := jwt.MapClaims{
		"id":  id,
		"exp": expirationTime.Unix(),
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)

	// Sign token with secret key
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenStr string) (string, error) {
	// Parse token
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Check signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	// Get data
	if data, ok := token.Claims.(jwt.MapClaims); ok {
		id := data["id"].(string)
		return id, nil
	}

	return "", fmt.Errorf("invalid claims")
}
