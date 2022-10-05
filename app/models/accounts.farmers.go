package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FarmersPendingResponse struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	AccountRequest AccountRequest     `json:"accountRequest" bson:"accountRequest"`
	Status         int                `json:"status" bson:"status"`
	Buyer          Employee           `json:"buyer" bson:"buyer"`
	Team           Team               `json:"team" bson:"team"`
	DenialReason   string             `json:"denialReason" bson:"denialReason"`
	DateCreated    int64              `json:"dateCreated" bson:"dateCreated"`
	DateUpdated    int64              `json:"dateUpdated" bson:"dateUpdated"`
	Description    string             `json:"description" bson:"description"`
	DownloadLink   string             `json:"downloadLink" bson:"downloadLink"`
}

type FarmersInworkResponse struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	AccountRequest AccountRequest     `json:"accountRequest" bson:"accountRequest"`
	Status         int                `json:"status" bson:"status"`
	Buyer          Employee           `json:"buyer" bson:"buyer"`
	Team           Team               `json:"team" bson:"team"`
	DenialReason   string             `json:"denialReason" bson:"denialReason"`
	DateCreated    int64              `json:"dateCreated" bson:"dateCreated"`
	DateUpdated    int64              `json:"dateUpdated" bson:"dateUpdated"`
	Description    string             `json:"description" bson:"description"`
	DownloadLink   string             `json:"downloadLink" bson:"downloadLink"`
}

type FarmersCompletedResponse struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	AccountRequest AccountRequest     `json:"accountRequest" bson:"accountRequest"`
	Status         int                `json:"status" bson:"status"`
	Buyer          Employee           `json:"buyer" bson:"buyer"`
	Team           Team               `json:"team" bson:"team"`
	Price          float64            `json:"price" bson:"price"`
	Valid          int                `json:"valid" bson:"valid"`
	Total          float64            `json:"totalSum" bson:"totalSum"`
	DenialReason   string             `json:"denialReason" bson:"denialReason"`
	DateCreated    int64              `json:"dateCreated" bson:"dateCreated"`
	DateFinished   int64              `json:"dateFinished" bson:"dateFinished"`
	Description    string             `json:"description" bson:"description"`
	DownloadLink   string             `json:"downloadLink" bson:"downloadLink"`
}

type FarmersCancelledResponse struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	AccountRequest AccountRequest     `json:"accountRequest" bson:"accountRequest"`
	Status         int                `json:"status" bson:"status"`
	Buyer          Employee           `json:"buyer" bson:"buyer"`
	CancelledBy    Employee           `json:"cancelledBy" bson:"cancelledBy"`
	Team           Team               `json:"team" bson:"team"`
	DenialReason   string             `json:"denialReason" bson:"denialReason"`
	DateCreated    int64              `json:"dateCreated" bson:"dateCreated"`
	DateCancelled  int64              `json:"dateCancelled" bson:"dateCancelled"`
	Description    string             `json:"description" bson:"description"`
	DownloadLink   string             `json:"downloadLink" bson:"downloadLink"`
}

type GroupedFarmersResponse struct {
	Farmer       Employee     `json:"_id" bson:"_id"`
	Teams        string       `json:"teams" bson:"teams"`
	Quantity     int          `json:"quantity" bson:"quantity"`
	Price        float64      `json:"price" bson:"price"`
	Valid        int          `json:"valid" bson:"valid"`
	Total        float64      `json:"totalSum" bson:"totalSum"`
	UserIdentity UserIdentity `json:"userIdentity"`
}
