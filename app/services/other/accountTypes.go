package services

import (
	"gypsylandFarming/app/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccountTypesService interface {
	CreateAccountType(*models.AccountType) error
	GetAll() ([]*models.AccountType, error)
	GetType(primitive.ObjectID) (*models.AccountType, error)
	GetTypeByName(string) (*models.AccountType, error)

	DeleteAll() (int, error)
	DeleteType(*primitive.ObjectID) error
}
