package models

import "go.mongodb.org/mongo-driver/bson/primitive"

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
	AccountRequest AccountRequest     `json:"accountRequest" bson:"accountRequest"`
	Status         int                `json:"status" bson:"status"`
	Buyer          Employee           `json:"buyer" bson:"buyer"`
	Farmer         Employee           `json:"farmer" bson:"farmer"`
	Team           Team               `json:"team" bson:"team"`
	Valid          int                `json:"valid" bson:"valid"`
	Price          float64            `json:"price" bson:"price"`
	TotalSum       float64            `json:"totalSum" bson:"totalSum"`
	Currency       Currency           `json:"currency" bson:"currency"`
	CancelledBy    Employee           `json:"cancelledBy" bson:"cancelledBy"`
	DateCreated    int64              `json:"dateCreated" bson:"dateCreated"`
	DateUpdated    int64              `json:"dateUpdated" bson:"dateUpdated"`
	DateFinished   int64              `json:"dateFinished" bson:"dateFinished"`
	DateCancelled  int64              `json:"dateCancelled" bson:"dateCancelled"`
	Description    string             `json:"description" bson:"description"`
	DenialReason   string             `json:"denialReason" bson:"denialReason"`
	DownloadLink   string             `json:"downloadLink" bson:"downloadLink"`
	FileName       string             `json:"fileName" bson:"fileName"`
}

type GetRequestBody struct {
	Period       Period       `json:"period"`
	UserIdentity UserIdentity `json:"userIdentity"`
	UserData     UserData     `json:"userData"`
	Status       int
}
