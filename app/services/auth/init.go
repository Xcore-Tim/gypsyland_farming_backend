package services

import (
	auth "gypsylandFarming/app/models/authentication"

	"github.com/golang-jwt/jwt/v4"
)

type AuthService interface {
	Login(*auth.UserCredentials, *auth.AuthResponseData) error

	GetFullname(*auth.Username, *auth.UserData) error
	AuthError(*auth.AuthResponseData, string)
}

type JWTService interface {
	GenerateToken(email string, isUser bool) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}
