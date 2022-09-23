package services

import (
	"github.com/golang-jwt/jwt/v4"
)

type AuthService interface {
}

type JWTService interface {
	GenerateToken(email string, isUser bool) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}
