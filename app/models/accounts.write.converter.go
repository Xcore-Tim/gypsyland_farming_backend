package models

import (
	"strconv"

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

	locationID, err := primitive.ObjectIDFromHex(r.AccountRequestBody.LocationID)

	if err != nil {
		panic("error parsing locationID")
	}

	typeID, err := primitive.ObjectIDFromHex(r.AccountRequestBody.TypeID)

	if err != nil {
		panic("error parsing locationID")
	}

	r.AccountRequestData.LocationID = locationID
	r.AccountRequestData.TypeID = typeID
	r.AccountRequestData.Quantity, err = strconv.Atoi(r.AccountRequestBody.Quantity)

	if err != nil {
		r.AccountRequestData.Quantity = 0
	}
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
