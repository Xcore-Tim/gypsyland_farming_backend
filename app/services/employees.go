package services

import (
	"gypsyland_farming/app/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmployeeService interface {
	CreateEmployee(*models.Employee) error
	GetEmployee(*primitive.ObjectID) (*models.Employee, error)
	GetAll() ([]*models.Employee, error)
	UpdateEmployee(*models.Employee) error
	DeleteEmployee(*primitive.ObjectID) error
}
