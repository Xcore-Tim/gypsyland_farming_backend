package services

import (
	"gypsyland_farming/app/models"

	"github.com/golang-jwt/jwt/v4"
)

type AuthService interface {
	CreateUser(*models.User) error
	UpdateUser(*models.User) error
	GetUser(authRequest *models.AuthRequest) (*models.User, bool)
	GetAll() ([]*models.User, error)
}

type JWTService interface {
	GenerateToken(email string, isUser bool) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}
