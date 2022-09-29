package services

import (
	"gypsyland_farming/app/models"

	"github.com/golang-jwt/jwt/v4"
)

type AuthService interface {
	Login(*models.UserCredentials, *models.AuthResponseData) error

	GetFullname(*models.Username, *models.UserData) error
	AuthError(*models.AuthResponseData, string)
}

type JWTService interface {
	GenerateToken(email string, isUser bool) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}
