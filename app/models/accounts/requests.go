package models

import (
	global "gypsylandFarming/app/models"
	auth "gypsylandFarming/app/models/authentication"
	currency "gypsylandFarming/app/models/currency"
	teams "gypsylandFarming/app/models/teams"

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

type AccountRequest struct {
	Type     global.AccountType `json:"type" bson:"type"`
	Location global.Location    `json:"location" bson:"location"`
	Quantity int                `json:"quantity" bson:"quantity"`
}

type AccountRequestTask struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	AccountRequest AccountRequest     `json:"accountRequest" bson:"accountRequest"`
	Status         int                `json:"status" bson:"status"`
	Buyer          global.Employee    `json:"buyer" bson:"buyer"`
	Farmer         global.Employee    `json:"farmer" bson:"farmer"`
	Team           teams.Team         `json:"team" bson:"team"`
	Valid          int                `json:"valid" bson:"valid"`
	Price          float64            `json:"price" bson:"price"`
	TotalSum       float64            `json:"totalSum" bson:"totalSum"`
	BaseTotal      float64            `json:"baseTotal" bson:"baseTotal"`
	Currency       currency.Currency  `json:"currency" bson:"currency"`
	BaseCurrency   currency.Currency  `json:"baseCurrency" bson:"baseCurrency"`
	CancelledBy    global.Employee    `json:"cancelledBy" bson:"cancelledBy"`
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
	Period       global.Period     `json:"period"`
	UserIdentity auth.UserIdentity `json:"userIdentity"`
	UserData     auth.UserData     `json:"userData"`
	Status       int
}
