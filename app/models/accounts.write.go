package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccountRequestBody struct {
	TypeID      string `json:"typeId"`
	LocationID  string `json:"locationId"`
	Quantity    string `json:"quantity"`
	Description string `json:"description"`
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

type CompleteAccountRequest struct {
	OrderID      string         `json:"orderID"`
	UserIdentity UserIdentity   `json:"userIdentity"`
	OrderInfo    CloseOrderInfo `json:"closeOrderInfo"`
	TotalSum     float64
	UserData     UserData
	RequestID    primitive.ObjectID
}

type CreateAccountRequestBody struct {
	UserIdentity   UserIdentity       `json:"userIdentity"`
	AccountRequest AccountRequestBody `json:"accountRequest"`
	UserData       UserData           `json:"userData,omitempty"`
}

type TakeAccountRequest struct {
	OrderID      string       `json:"orderID"`
	UserIdentity UserIdentity `json:"userIdentity"`
	RequestID    primitive.ObjectID
	UserData     UserData
	Farmer       Employee
}

// type UpdateRequestBody struct {
// 	OrderID        string             `json:"orderID"`
// 	UserIdentity   UserIdentity       `json:"userIdentity"`
// 	UpdateBody     AccountRequestBody `json:"updateBody"`
// 	RequestID      primitive.ObjectID
// 	UserData       UserData
// 	UpdateData     UpdateRequestData
// 	AccountRequest AccountRequest
// }

type UpdateRequestBody struct {
	OrderID        string            `json:"orderID"`
	UserIdentity   UserIdentity      `json:"userIdentity"`
	UpdateBody     UpdateRequestData `json:"updateBody"`
	RequestID      primitive.ObjectID
	UserData       UserData
	AccountRequest AccountRequest
}

// type UpdateRequestData struct {
// 	TypeID     primitive.ObjectID
// 	LocationID primitive.ObjectID
// 	Quantity   int
// }

type UpdateRequestData struct {
	AccountType string `json:"accountType"`
	Location    string `json:"location"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
