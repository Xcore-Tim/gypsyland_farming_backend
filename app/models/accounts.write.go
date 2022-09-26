package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TakeAccountRequest struct {
	OrderID      string       `json:"orderID"`
	UserIdentity UserIdentity `json:"userIdentity"`
	RequestID    primitive.ObjectID
	UserData     UserData
	Farmer       Employee
}

type CompleteAccountRequest struct {
	OrderID      string         `json:"orderID"`
	UserIdentity UserIdentity   `json:"userIdentity"`
	OrderInfo    CloseOrderInfo `json:"closeOrderInfo"`
	TotalSum     float64
	UserData     UserData
	RequestID    primitive.ObjectID
}

type CancelAccountRequest struct {
	OrderID      string       `json:"orderID"`
	UserIdentity UserIdentity `json:"userIdentity"`
	DenialReason string       `json:"denialReason"`
	UserData     UserData
	RequestID    primitive.ObjectID
	CancelledBy  Employee
}

type CloseOrderInfo struct {
	Valid       int     `json:"valid"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Link        string  `json:"downloadLink"`
}

type UpdateAccountRequest struct {
	OrderID        string         `json:"orderID"`
	UserIdentity   UserIdentity   `json:"userIdentity"`
	AccountRequest AccountRequest `json:"accountRequest"`
	Description    string         `json:"description"`
	RequestID      primitive.ObjectID
	UserData       UserData
}

type AccountRequestBody struct {
	TypeID      string `json:"typeId"`
	LocationID  string `json:"locationId"`
	Quantity    string `json:"quantity"`
	Description string `json:"description"`
}

type CreateAccountRequestBody struct {
	UserIdentity   UserIdentity       `json:"userIdentity"`
	AccountRequest AccountRequestBody `json:"accountRequest"`
	UserData       UserData           `json:"userData,omitempty"`
}
