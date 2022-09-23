package models

import (
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *CancelAccountRequest) Convert() {
	r.UserData.UserID, _ = strconv.Atoi(r.UserIdentity.UserID)
	r.UserData.TeamID, _ = strconv.Atoi(r.UserIdentity.TeamID)
	r.UserData.RoleID, _ = strconv.Atoi(r.UserIdentity.RoleID)
	r.UserData.Username = r.UserIdentity.Username
	r.UserData.Token = r.UserIdentity.Token
	r.RequestID, _ = primitive.ObjectIDFromHex(r.OrderID)
	r.CancelledBy = Employee{
		ID:       r.UserData.UserID,
		Name:     r.UserData.Username,
		Position: r.UserData.RoleID,
	}
}

func (r *CreateAccountRequestBody) Convert() {
	r.UserData.UserID, _ = strconv.Atoi(r.UserIdentity.UserID)
	r.UserData.TeamID, _ = strconv.Atoi(r.UserIdentity.TeamID)
	r.UserData.RoleID, _ = strconv.Atoi(r.UserIdentity.RoleID)
	r.UserData.Username = r.UserIdentity.Username
	r.UserData.Token = r.UserIdentity.Token
}

func (r *UpdateAccountRequest) Convert() {
	r.UserData.UserID, _ = strconv.Atoi(r.UserIdentity.UserID)
	r.UserData.TeamID, _ = strconv.Atoi(r.UserIdentity.TeamID)
	r.UserData.RoleID, _ = strconv.Atoi(r.UserIdentity.RoleID)
	r.UserData.Username = r.UserIdentity.Username
	r.UserData.Token = r.UserIdentity.Token
	r.RequestID, _ = primitive.ObjectIDFromHex(r.OrderID)
}

func (r *TakeAccountRequest) Convert() {
	r.UserData.UserID, _ = strconv.Atoi(r.UserIdentity.UserID)
	r.UserData.TeamID, _ = strconv.Atoi(r.UserIdentity.TeamID)
	r.UserData.RoleID, _ = strconv.Atoi(r.UserIdentity.RoleID)
	r.UserData.Username = r.UserIdentity.Username
	r.UserData.Token = r.UserIdentity.Token
	r.RequestID, _ = primitive.ObjectIDFromHex(r.OrderID)
	r.Farmer = Employee{
		ID:       r.UserData.UserID,
		Name:     r.UserData.Username,
		Position: r.UserData.RoleID,
	}
}

func (r *CompleteAccountRequest) Convert() {
	r.UserData.UserID, _ = strconv.Atoi(r.UserIdentity.UserID)
	r.UserData.TeamID, _ = strconv.Atoi(r.UserIdentity.TeamID)
	r.UserData.RoleID, _ = strconv.Atoi(r.UserIdentity.RoleID)
	r.UserData.Username = r.UserIdentity.Username
	r.UserData.Token = r.UserIdentity.Token
	r.RequestID, _ = primitive.ObjectIDFromHex(r.OrderID)
}
