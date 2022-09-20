package services

import (
	"gypsyland_farming/app/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PositionService interface {
	CreatePosition(*models.Position) error
	GetPosition(*primitive.ObjectID) (*models.Position, error)
	GetAll() ([]*models.Position, error)
	UpdatePosition(*models.Position) error
	DeletePosition(*primitive.ObjectID) error
}
