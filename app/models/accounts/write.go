package models

import (
	global "gypsylandFarming/app/models"
	auth "gypsylandFarming/app/models/authentication"
	teams "gypsylandFarming/app/models/teams"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccountRequestBody struct {
	TypeID      string `json:"typeId"`
	LocationID  string `json:"locationId"`
	Quantity    string `json:"quantity"`
	Price       string `json:"price"`
	Description string `json:"description"`
}

type AccountRequestData struct {
	TypeID     primitive.ObjectID
	LocationID primitive.ObjectID
	Quantity   int
	Price      float64
}

type CancelAccountRequest struct {
	OrderID      string            `json:"orderID"`
	UserIdentity auth.UserIdentity `json:"userIdentity"`
	DenialReason string            `json:"denialReason"`
	UserData     auth.UserData
	RequestID    primitive.ObjectID
	CancelledBy  global.Employee
}

type CreateAccountRequestBody struct {
	UserIdentity       auth.UserIdentity  `json:"userIdentity"`
	AccountRequestBody AccountRequestBody `json:"accountRequest"`
	UserData           auth.UserData      `json:"userData,omitempty"`
	AccountRequestData AccountRequestData
}

type CloseOrderInfo struct {
	Valid       int     `json:"valid"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Link        string  `json:"downloadLink"`
}

type CompleteAccountRequest struct {
	OrderID      string            `json:"orderID"`
	UserIdentity auth.UserIdentity `json:"userIdentity"`
	OrderInfo    CloseOrderInfo    `json:"closeOrderInfo"`
	TotalSum     float64
	UserData     auth.UserData
	RequestID    primitive.ObjectID
}

type TakeAccountRequest struct {
	OrderID      string            `json:"orderID"`
	UserIdentity auth.UserIdentity `json:"userIdentity"`
	RequestID    primitive.ObjectID
	UserData     auth.UserData
	Farmer       global.Employee
}

type UpdateRequestBody struct {
	UserIdentity auth.UserIdentity `json:"userIdentity"`
	UpdateBody   UpdateRequestData `json:"updateBody"`
	RequestID    primitive.ObjectID
	UserData     auth.UserData
}

type UCResponseBody struct {
	ID             primitive.ObjectID
	AccountRequest AccountRequest  `json:"accountRequest" bson:"accountRequest"`
	Buyer          global.Employee `json:"buyer"`
	Team           teams.Team      `json:"team"`
	Valid          int             `json:"valid"`
	Price          float64         `json:"price"`
	Total          float64         `json:"totalSum"`
	Description    string          `jspn:"description"`
}

type UpdateRequestData struct {
	AccountType string  `json:"accountType"`
	Location    string  `json:"location"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}
