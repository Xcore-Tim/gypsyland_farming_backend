package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Employee struct {
	ID       int    `json:"id" bson:"id"`
	Name     string `json:"name" bson:"name"`
	Position int    `json:"position" bson:"position"`
}

type Location struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name"`
	ISO  string             `json:"iso" bson:"iso"`
}

type Position struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string             `json:"name" bson:"name"`
	RoleNumber int                `json:"role_number" bson:"role_number"`
}

type Repository struct {
	DiskLocation string
	Link         string
}

type Currency struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name"`
	ISO  string             `json:"iso" bson:"iso"`
}

type Period struct {
	StartISO  string    `json:"startDate"`
	EndISO    string    `json:"endDate"`
	StartDate time.Time `json:"unixStart"`
	EndDate   time.Time `json:"unixEnd"`
}

type AccountType struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name"`
}
