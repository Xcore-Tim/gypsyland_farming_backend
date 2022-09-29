package services

import (
	"gypsyland_farming/app/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LocationService interface {
	CreateLocation(*models.Location) error
	GetLocation(primitive.ObjectID) (*models.Location, error)
	GetLocationByName(string) (*models.Location, error)
	GetAll() ([]*models.Location, error)
	UpdateLocation(*models.Location) error
	DeleteLocation(*string) error
}
