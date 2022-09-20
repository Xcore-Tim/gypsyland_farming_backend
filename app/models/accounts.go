package models

import (
	// "time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	Pending  = 0
	Inwork   = 1
	Complete = 2
	Canceled = 3
	Returned = 4

	Admin          = 1
	TeamLead       = 2
	Buyer          = 3
	Smart          = 4
	TeamLeadFarmer = 5
	Farmer         = 6
	Helper         = 7
	Creative       = 8
)

type AccountType struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name"`
}

type AccountRequest struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Type     AccountType        `json:"type" bson:"type"`
	Location Location           `json:"location" bson:"location"`
	Quantity int                `json:"quantity" bson:"quantity"`
}

type AccountRequestTask struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	AccountRequest AccountRequest     `json:"account_request" bson:"account_request"`
	Status         int                `json:"status" bson:"status"`
	Buyer          Employee           `json:"buyer" bson:"buyer"`
	BuyerID        int                `json:"buyerID" bson:"buyerID"`
	Farmer         Employee           `json:"farmer" bson:"farmer"`
	FarmerID       int                `json:"farmerID" bson:"farmerID"`
	Team           Team               `json:"team" bson:"team"`
	TeamID         int                `json:"teamID" bson:"teamID"`
	Valid          int                `json:"valid" bson:"valid"`
	Price          float64            `json:"price" bson:"price"`
	Currency       Currency           `json:"currency" bson:"currency"`
	DenialReason   string             `json:"denial_reason" bson:"denial_reason"`
	DateCreated    int64              `json:"date_created" bson:"date_created"`
	DateUpdated    int64              `json:"date_updated" bson:"date_updated"`
	Description    string             `json:"description" bson:"description"`
	TotalSum       float64            `json:"total_sum" bson:"total_sum"`
}

type AccountRequestCompleted struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Valid       int                `json:"valid" bson:"valid"`
	Price       float64            `json:"price" bson:"price"`
	Description string             `json:"description" bson:"description"`
	TotalSum    float64            `json:"total_sum" bson:"total_sum"`
}

type AccountRequestCanceled struct {
	Description string `json:"description"`
}

type AccountRequestUpdate struct {
	ID           primitive.ObjectID
	Quantity     int     `json:"quantity" bson:"quantity"`
	Price        float64 `json:"price" bson:"price"`
	Description  string  `json:"description" bson:"description"`
	DenialReason string  `json:"denial_reason" bson:"denial_reason"`
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

type GetRequestBody struct {
	Period       Period       `json:"period"`
	UserIdentity UserIdentity `json:"user_identity"`
	Status       int
}

type PostRequestBody struct {
	UserIdentity   UserIdentity   `json:"user_identity"`
	AccountRequest AccountRequest `json:"account_request"`
	Description    string         `json:"description"`
}
