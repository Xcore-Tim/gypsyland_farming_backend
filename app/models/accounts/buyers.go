package models

import (
	global "gypsylandFarming/app/models"
	teams "gypsylandFarming/app/models/teams"

	auth "gypsylandFarming/app/models/authentication"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BuyersPendingResponse struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	AccountRequest AccountRequest     `json:"accountRequest" bson:"accountRequest"`
	Status         int                `json:"status" bson:"status"`
	Farmer         global.Employee    `json:"farmer" bson:"farmer"`
	Team           teams.Team         `json:"team" bson:"team"`
	DenialReason   string             `json:"denialReason" bson:"denialReason"`
	DateCreated    int64              `json:"dateCreated" bson:"dateCreated"`
	DateUpdated    int64              `json:"dateUpdated" bson:"dateUpdated"`
	Description    string             `json:"description" bson:"description"`
	DownloadLink   string             `json:"downloadLink" bson:"downloadLink"`
}

type BuyersInworkResponse struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	AccountRequest AccountRequest     `json:"accountRequest" bson:"accountRequest"`
	Status         int                `json:"status" bson:"status"`
	Farmer         global.Employee    `json:"farmer" bson:"farmer"`
	Team           teams.Team         `json:"team" bson:"team"`
	DenialReason   string             `json:"denialReason" bson:"denialReason"`
	DateCreated    int64              `json:"dateCreated" bson:"dateCreated"`
	DateUpdated    int64              `json:"dateUpdated" bson:"dateUpdated"`
	Description    string             `json:"description" bson:"description"`
	DownloadLink   string             `json:"downloadLink" bson:"downloadLink"`
}

type BuyersCompletedResponse struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	AccountRequest AccountRequest     `json:"accountRequest" bson:"accountRequest"`
	Status         int                `json:"status" bson:"status"`
	Farmer         global.Employee    `json:"farmer" bson:"farmer"`
	Team           teams.Team         `json:"team" bson:"team"`
	Price          float64            `json:"price" bson:"price"`
	Valid          int                `json:"valid" bson:"valid"`
	Total          float64            `json:"totalSum" bson:"totalSum"`
	DenialReason   string             `json:"denialReason" bson:"denialReason"`
	DateCreated    int64              `json:"dateCreated" bson:"dateCreated"`
	DateFinished   int64              `json:"dateFinished" bson:"dateFinished"`
	Description    string             `json:"description" bson:"description"`
	DownloadLink   string             `json:"downloadLink" bson:"downloadLink"`
}

type BuyersCancelledResponse struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	AccountRequest AccountRequest     `json:"accountRequest" bson:"accountRequest"`
	Status         int                `json:"status" bson:"status"`
	Farmer         global.Employee    `json:"farmer" bson:"farmer"`
	CancelledBy    global.Employee    `json:"cancelledBy" bson:"cancelledBy"`
	Team           teams.Team         `json:"team" bson:"team"`
	DenialReason   string             `json:"denialReason" bson:"denialReason"`
	DateCreated    int64              `json:"dateCreated" bson:"dateCreated"`
	DateCancelled  int64              `json:"dateCancelled" bson:"dateCancelled"`
	Description    string             `json:"description" bson:"description"`
	DownloadLink   string             `json:"downloadLink" bson:"downloadLink"`
}

type GroupedBuyersResponse struct {
	Buyer    global.Employee `json:"_id" bson:"_id"`
	Team     teams.Team      `json:"team" bson:"team"`
	Quantity int             `json:"quantity" bson:"quantity"`
	Valid    int             `json:"valid" bson:"valid"`
	Price    float64         `json:"price" bson:"price"`
	Total    float64         `json:"totalSum" bson:"totalSum"`
	UserData auth.UserData   `json:"userData"`
}

type GroupedTeamsResponse struct {
	ID       teams.Team `json:"_id" bson:"_id"`
	Quantity int        `json:"quantity" bson:"quantity"`
	Price    float64    `json:"price" bson:"price"`
	Valid    int        `json:"valid" bson:"valid"`
	Total    float64    `json:"totalSum" bson:"totalSum"`
}
