package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(claims jwt.MapClaims, method jwt.SigningMethod, secret string) (string, error) {
	return jwt.NewWithClaims(method, claims).SignedString([]byte(secret))
}

func ValidateToken(token string, secret string) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	return parsedToken, nil
}
