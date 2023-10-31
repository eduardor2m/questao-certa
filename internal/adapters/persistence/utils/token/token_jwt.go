package token

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
)

func StringToJWT(tokenReceived string) (*jwt.Token, error) {
	jwtSecretKey := os.Getenv("JWT_SECRET")

	token, err := jwt.Parse(tokenReceived, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("falha ao verificar token: %v", err)
	}

	return token, nil
}

func ExtractClainsFromJwtToken(token *jwt.Token) (jwt.MapClaims, error) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("token inválido")
	}
}
