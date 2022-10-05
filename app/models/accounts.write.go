package models

import (
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
	OrderID      string       `json:"orderID"`
	UserIdentity UserIdentity `json:"userIdentity"`
	DenialReason string       `json:"denialReason"`
	UserData     UserData
	RequestID    primitive.ObjectID
	CancelledBy  Employee
}

type CreateAccountRequestBody struct {
	UserIdentity       UserIdentity       `json:"userIdentity"`
	AccountRequestBody AccountRequestBody `json:"accountRequest"`
	UserData           UserData           `json:"userData,omitempty"`
	AccountRequestData AccountRequestData
}

type CloseOrderInfo struct {
	Valid       int     `json:"valid"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Link        string  `json:"downloadLink"`
}

type CompleteAccountRequest struct {
	OrderID      string         `json:"orderID"`
	UserIdentity UserIdentity   `json:"userIdentity"`
	OrderInfo    CloseOrderInfo `json:"closeOrderInfo"`
	TotalSum     float64
	UserData     UserData
	RequestID    primitive.ObjectID
}

type TakeAccountRequest struct {
	OrderID      string       `json:"orderID"`
	UserIdentity UserIdentity `json:"userIdentity"`
	RequestID    primitive.ObjectID
	UserData     UserData
	Farmer       Employee
}

type UpdateRequestBody struct {
	OrderID        string            `json:"orderID"`
	UserIdentity   UserIdentity      `json:"userIdentity"`
	UpdateBody     UpdateRequestData `json:"updateBody"`
	RequestID      primitive.ObjectID
	UserData       UserData
	AccountRequest AccountRequest
}

type UpdateRequestData struct {
	AccountType string `json:"accountType"`
	Location    string `json:"location"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
