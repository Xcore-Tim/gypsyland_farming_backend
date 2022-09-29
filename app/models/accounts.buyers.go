package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetFunctions func(requestBody *GetRequestBody) bson.D

type BuyersPendingResponse struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	AccountRequest AccountRequest     `json:"accountRequest" bson:"accountRequest"`
	Status         int                `json:"status" bson:"status"`
	Farmer         Employee           `json:"farmer" bson:"farmer"`
	Team           Team               `json:"team" bson:"team"`
	DenialReason   string             `json:"denialReason" bson:"denialReason"`
	DateCreated    int64              `json:"dateCreated" bson:"dateCreated"`
	DateUpdated    int64              `json:"dateUpdated" bson:"dateUpdated"`
	Description    string             `json:"description" bson:"description"`
}

type BuyersInworkResponse struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	AccountRequest AccountRequest     `json:"accountRequest" bson:"accountRequest"`
	Status         int                `json:"status" bson:"status"`
	Farmer         Employee           `json:"farmer" bson:"farmer"`
	Team           Team               `json:"team" bson:"team"`
	DenialReason   string             `json:"denialReason" bson:"denialReason"`
	DateCreated    int64              `json:"dateCreated" bson:"dateCreated"`
	DateUpdated    int64              `json:"dateUpdated" bson:"dateUpdated"`
	Description    string             `json:"description" bson:"description"`
}

type BuyersCompletedResponse struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	AccountRequest AccountRequest     `json:"accountRequest" bson:"accountRequest"`
	Status         int                `json:"status" bson:"status"`
	Farmer         Employee           `json:"farmer" bson:"farmer"`
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

type BuyersCancelledResponse struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	AccountRequest AccountRequest     `json:"accountRequest" bson:"accountRequest"`
	Status         int                `json:"status" bson:"status"`
	Farmer         Employee           `json:"farmer" bson:"farmer"`
	Team           Team               `json:"team" bson:"team"`
	DenialReason   string             `json:"denialReason" bson:"denialReason"`
	DateCreated    int64              `json:"dateCreated" bson:"dateCreated"`
	DateCancelled  int64              `json:"dateCancelled" bson:"dateCancelled"`
	Description    string             `json:"description" bson:"description"`
}

type GroupedBuyersResponse struct {
	Buyer    Employee `json:"_id" bson:"_id"`
	Team     Team     `json:"team" bson:"team"`
	Quantity int      `json:"quantity" bson:"quantity"`
	Valid    int      `json:"valid" bson:"valid"`
	Price    float64  `json:"price" bson:"price"`
	Total    float64  `json:"totalSum" bson:"totalSum"`
	UserData UserData `json:"userData"`
}

type GroupedTeamsResponse struct {
	ID       Team    `json:"_id" bson:"_id"`
	Quantity int     `json:"quantity" bson:"quantity"`
	Price    float64 `json:"price" bson:"price"`
	Valid    int     `json:"valid" bson:"valid"`
	Total    float64 `json:"totalSum" bson:"totalSum"`
}
