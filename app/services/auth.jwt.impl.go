package services

import (
	"errors"
	"gypsyland_farming/app/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTServicesImpl struct {
	// customClaims *models.AuthCustomClaims
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &JWTServicesImpl{
		secretKey: getSecretKey(),
		issuer:    "GypsyServer",
	}
}

func getSecretKey() string {
	secret := os.Getenv("SECRET")

	if secret == "" {
		secret = "secret"
	}

	return secret
}

func (srv *JWTServicesImpl) GenerateToken(email string, isUser bool) (string, error) {

	// srv.customClaims.Name = email
	// srv.customClaims.IsUser = isUser

	// srv.customClaims.StandardClaims.ExpiresAt = time.Now().Add(time.Hour * 2).Unix()
	// srv.customClaims.StandardClaims.Issuer = srv.issuer
	// srv.customClaims.StandardClaims.IssuedAt = time.Now().Unix()

	customClaims := &models.AuthCustomClaims{
		Name: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    srv.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)

	t, err := token.SignedString([]byte(srv.secretKey))

	if err != nil {
		return "", err
	}

	return t, err
}

func (srv *JWTServicesImpl) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, errors.New("invalid token")
		}
		return []byte(srv.secretKey), nil
	})
}
