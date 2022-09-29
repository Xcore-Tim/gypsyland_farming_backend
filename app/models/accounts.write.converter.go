package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *CancelAccountRequest) Convert() {

	ConvertUserData(&r.UserData, r.UserIdentity)

	r.RequestID, _ = primitive.ObjectIDFromHex(r.OrderID)
	r.CancelledBy = Employee{
		ID:       r.UserData.UserID,
		Name:     r.UserData.Username,
		Position: r.UserData.RoleID,
	}
}

func (r *CreateAccountRequestBody) Convert() {
	ConvertUserData(&r.UserData, r.UserIdentity)
}

func (r *UpdateRequestBody) Convert() {
	ConvertUserData(&r.UserData, r.UserIdentity)
	r.RequestID, _ = primitive.ObjectIDFromHex(r.OrderID)
}

func (r *TakeAccountRequest) Convert() {

	ConvertUserData(&r.UserData, r.UserIdentity)

	r.RequestID, _ = primitive.ObjectIDFromHex(r.OrderID)
	r.Farmer = Employee{
		ID:       r.UserData.UserID,
		Name:     r.UserData.Username,
		Position: r.UserData.RoleID,
	}
}

func (r *CompleteAccountRequest) Convert() {
	ConvertUserData(&r.UserData, r.UserIdentity)
	r.RequestID, _ = primitive.ObjectIDFromHex(r.OrderID)
}
