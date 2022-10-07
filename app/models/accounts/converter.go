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

	locationID, err := primitive.ObjectIDFromHex(model.AccountRequestBody.LocationID)

	if err != nil {
		panic("error parsing locationID")
	}

	typeID, err := primitive.ObjectIDFromHex(model.AccountRequestBody.TypeID)

	if err != nil {
		panic("error parsing locationID")
	}

	model.AccountRequestData.LocationID = locationID
	model.AccountRequestData.TypeID = typeID
	model.AccountRequestData.Quantity, err = strconv.Atoi(model.AccountRequestBody.Quantity)

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

	if period.StartISO == "" {
		period.EndDate = time.Now()
	} else if period.EndISO == "" {
		period.EndDate = time.Now()
	} else {
		date_format := "2006-01-02"
		period.EndDate, _ = time.Parse(date_format, period.EndISO)
		period.StartDate, _ = time.Parse(date_format, period.StartISO)
	}

}
