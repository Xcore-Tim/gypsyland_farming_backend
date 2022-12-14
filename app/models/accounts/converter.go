package models

import (
	global "gypsylandFarming/app/models"
	auth "gypsylandFarming/app/models/authentication"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (model *GetRequestBody) Convert() {
	ConvertUserData(&model.UserData, model.UserIdentity)
	ConvertPeriod(&model.Period)
}

func (model *CancelAccountRequest) Convert() {

	ConvertUserData(&model.UserData, model.UserIdentity)

	model.RequestID, _ = primitive.ObjectIDFromHex(model.OrderID)
	model.CancelledBy = global.Employee{
		ID:       model.UserData.UserID,
		Name:     model.UserData.Username,
		Position: model.UserData.RoleID,
	}
}

func (model *CreateAccountRequestBody) Convert() {

	ConvertUserData(&model.UserData, model.UserIdentity)

	if model.AccountRequestBody.LocationID != "" {
		locationID, _ := primitive.ObjectIDFromHex(model.AccountRequestBody.LocationID)
		model.AccountRequestData.LocationID = locationID
	}

	typeID, err := primitive.ObjectIDFromHex(model.AccountRequestBody.TypeID)

	if err != nil {
		panic("error parsing locationID")
	}

	model.AccountRequestData.TypeID = typeID
	model.AccountRequestData.Quantity, _ = strconv.Atoi(model.AccountRequestBody.Quantity)
	model.AccountRequestData.Price, _ = strconv.ParseFloat(model.AccountRequestBody.Price, 64)

	if err != nil {
		model.AccountRequestData.Quantity = 0
	}
}

func (model *UpdateRequestBody) Convert() {
	ConvertUserData(&model.UserData, model.UserIdentity)
}

func (model *TakeAccountRequest) Convert() {

	ConvertUserData(&model.UserData, model.UserIdentity)

	model.RequestID, _ = primitive.ObjectIDFromHex(model.OrderID)
	model.Farmer = global.Employee{
		ID:       model.UserData.UserID,
		Name:     model.UserData.Username,
		Position: model.UserData.RoleID,
	}
}

func (model *CompleteAccountRequest) Convert() {
	ConvertUserData(&model.UserData, model.UserIdentity)
	model.RequestID, _ = primitive.ObjectIDFromHex(model.OrderID)
}

func ConvertUserData(userData *auth.UserData, userIdentity auth.UserIdentity) {

	userData.UserID, _ = strconv.Atoi(userIdentity.UserID)
	userData.TeamID, _ = strconv.Atoi(userIdentity.TeamID)
	userData.RoleID, _ = strconv.Atoi(userIdentity.RoleID)
	userData.Username = userIdentity.Username
	userData.Token = userIdentity.Token
}

func ConvertPeriod(period *global.Period) {

	date_format := "2006-01-02"

	if period.StartISO == "" {
		period.StartISO = "1970-01-01"
	}

	period.StartDate, _ = time.Parse(date_format, period.StartISO)

	if period.EndISO != "" {
		period.EndDate, _ = time.Parse(date_format, period.EndISO)
		return
	}

	period.EndDate = time.Now()
}
