package auth

import (
	"time"
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("YOUR_SECRET_KEY") // load from env

func GenerateJWT(employeeID string, hours int) (string, error) {
	claims := jwt.MapClaims{}
	claims["employee_id"] = employeeID
	claims["exp"] = time.Now().Add(time.Duration(hours) * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
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