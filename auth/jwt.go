package auth

import (
	"time"
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("YOUR_SECRET_KEY") // load from env

type TokenType string

const (
	AccessToken  TokenType = "access"
	RefreshToken TokenType = "refresh"
)

func GenerateToken(employeeID string, duration time.Duration, tokenType TokenType) (string, error) {

	claims := jwt.MapClaims{
		"employee_id": employeeID,
		"exp":         time.Now().Add(duration).Unix(),
		"type":        tokenType,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func GenerateAccessToken(employeeID string, hours int) (string, error) {
	return GenerateToken(employeeID, time.Duration(hours)*time.Hour, AccessToken)
}

func GenerateRefreshToken(employeeID string, days int) (string, error) {
	return GenerateToken(employeeID, time.Duration(days)*24*time.Hour, RefreshToken)
}

func ValidateJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid claims")
	}

	employeeID, ok := claims["employee_id"].(string)
	if !ok {
		return "", errors.New("employee_id not found in token")
	}

	return employeeID, nil
}