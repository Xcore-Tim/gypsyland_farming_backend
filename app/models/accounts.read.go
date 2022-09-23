package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetFunctions func(requestBody *GetRequestBody) bson.D

type GetRequestBody struct {
	Period       Period       `json:"period"`
	UserIdentity UserIdentity `json:"userIdentity"`
	UserData     UserData     `json:"userData"`
	Status       int
}

type BuyersPendingResponse struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	AccountRequest AccountRequest     `json:"accountRequest" bson:"accountRequest"`
	Status         int                `json:"status" bson:"status"`
	Buyer          Employee           `json:"buyer" bson:"buyer"`
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
	Buyer          Employee           `json:"buyer" bson:"buyer"`
	Team           Team               `json:"team" bson:"team"`
	DenialReason   string             `json:"denialReason" bson:"denialReason"`
	DateCreated    int64              `json:"dateCreated" bson:"dateCreated"`
	DateCancelled  int64              `json:"dateCancelled" bson:"dateCancelled"`
	Description    string             `json:"description" bson:"description"`
}

type GroupedFarmersResponse struct {
	Farmer   Employee `json:"_id" bson:"_id"`
	Teams    string   `json:"teams" bson:"teams"`
	Quantity int      `json:"quantity" bson:"quantity"`
	Price    float64  `json:"price" bson:"price"`
	Valid    int      `json:"valid" bson:"valid"`
}

type GroupedBuyersResponse struct {
	Buyer    Employee `json:"_id" bson:"_id"`
	Team     Team     `josn:"team" bson:"team"`
	Quantity int      `json:"quantity" bson:"quantity"`
	Price    float64  `json:"price" bson:"price"`
	Valid    int      `json:"valid" bson:"valid"`
}

type GroupedTeamsResponse struct {
	ID       Team    `json:"_id" bson:"_id"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
	Valid    int     `json:"valid"`
}
