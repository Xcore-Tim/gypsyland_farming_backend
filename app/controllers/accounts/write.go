package controllers

import (
	accounts "gypsylandFarming/app/models/accounts"

	"time"

	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ctrl AccountRequestController) Test(ctx *gin.Context) {

	oid := ctx.Query("oid")
	ctx.JSON(http.StatusOK, oid)
}

func (ctrl AccountRequestController) CreateAccountRequest(ctx *gin.Context) {

	var requestBody accounts.CreateAccountRequestBody

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}

	requestBody.Convert()

	var accountRequestTask accounts.AccountRequestTask

	if requestBody.AccountRequestBody.LocationID != "" {
		location, err := ctrl.LocationService.GetLocation(requestBody.AccountRequestData.LocationID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			// return
		}
		accountRequestTask.AccountRequest.Location = *location
	}

	accountType, err := ctrl.AccountTypesService.GetType(requestBody.AccountRequestData.TypeID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	team, err := ctrl.TeamService.GetTeamByNum(requestBody.UserData.TeamID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accountRequestTask.AccountRequest.Type = *accountType
	accountRequestTask.AccountRequest.Quantity = requestBody.AccountRequestData.Quantity

	accountRequestTask.DateCreated = time.Now().Unix()
	accountRequestTask.Description = requestBody.AccountRequestBody.Description
	accountRequestTask.Price = requestBody.AccountRequestData.Price
	total := float64(accountRequestTask.AccountRequest.Quantity) * accountRequestTask.Price
	accountRequestTask.TotalSum = ctrl.WriteAccountRequestService.RoundFloat(total, 2)

	accountRequestTask.Team = *team
	accountRequestTask.Buyer.ID = requestBody.UserData.UserID
	accountRequestTask.Buyer.Name = requestBody.UserData.Username
	accountRequestTask.Buyer.Position = requestBody.UserData.RoleID

	if err := ctrl.WriteAccountRequestService.CreateAccountRequest(&accountRequestTask); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": accountRequestTask.ID})
}

func (ctrl AccountRequestController) UpdateRequest(ctx *gin.Context) {

	orderID, err := primitive.ObjectIDFromHex(ctx.Query("orderID"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	var accountRequestUpdate accounts.UpdateRequestBody

	if err := ctx.ShouldBindJSON(&accountRequestUpdate); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	accountRequestUpdate.RequestID = orderID
	accountRequestUpdate.Convert()

	originalAccountRequest, err := ctrl.ReadAccountRequestService.GetAccountRequestData(&accountRequestUpdate.RequestID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var updateResponse accounts.UCResponseBody

	updateResponse.ID = originalAccountRequest.ID
	updateResponse.AccountRequest = originalAccountRequest.AccountRequest
	updateResponse.Description = accountRequestUpdate.UpdateBody.Description

	if accountRequestUpdate.UpdateBody.Location != "" {
		if location, err := ctrl.LocationService.GetLocationByName(accountRequestUpdate.UpdateBody.Location); err == nil {
			updateResponse.AccountRequest.Location = *location
		}
	}

	if accountRequestUpdate.UpdateBody.AccountType != "" {
		if accountType, err := ctrl.AccountTypesService.GetTypeByName(accountRequestUpdate.UpdateBody.AccountType); err == nil {
			updateResponse.AccountRequest.Type = *accountType
		}
	}

	if accountRequestUpdate.UpdateBody.Quantity != originalAccountRequest.AccountRequest.Quantity {
		updateResponse.AccountRequest.Quantity = accountRequestUpdate.UpdateBody.Quantity
	}

	if accountRequestUpdate.UpdateBody.Price != originalAccountRequest.Price {
		updateResponse.Price = accountRequestUpdate.UpdateBody.Price
	}

	total := float64(updateResponse.AccountRequest.Quantity) * updateResponse.Price
	updateResponse.Total = ctrl.WriteAccountRequestService.RoundFloat(total, 2)

	ctx.JSON(http.StatusOK, updateResponse)

	if err := ctrl.WriteAccountRequestService.UpdateRequest(&updateResponse); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (ctrl AccountRequestController) TakeAccountRequest(ctx *gin.Context) {

	var requestData accounts.TakeAccountRequest

	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	requestData.Convert()

	if err := ctrl.WriteAccountRequestService.TakeAccountRequest(&requestData); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (ctrl AccountRequestController) CancelAccountRequest(ctx *gin.Context) {

	var cancelRequest accounts.CancelAccountRequest

	err := ctx.ShouldBindJSON(&cancelRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	cancelRequest.Convert()

	err = ctrl.WriteAccountRequestService.CancelAccountRequest(&cancelRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (ctrl AccountRequestController) CompleteAccountRequest(ctx *gin.Context) {

	var accountRequestCompleted accounts.CompleteAccountRequest

	if err := ctx.ShouldBindJSON(&accountRequestCompleted); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	accountRequestCompleted.Convert()

	accountRequest, err := ctrl.ReadAccountRequestService.GetRequestTask(&accountRequestCompleted.RequestID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}

	if accountRequest.Price != accountRequestCompleted.OrderInfo.Price {
		total := float64(accountRequest.AccountRequest.Quantity) * accountRequestCompleted.OrderInfo.Price
		accountRequestCompleted.TotalSum = ctrl.WriteAccountRequestService.RoundFloat(total, 2)
	}

	if err := ctrl.WriteAccountRequestService.CompleteAccountRequest(&accountRequestCompleted); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (ctrl AccountRequestController) ReturnAccountRequest(ctx *gin.Context) {

	var requestData accounts.TakeAccountRequest

	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	requestData.Convert()

	if err := ctrl.WriteAccountRequestService.ReturnAccountRequest(&requestData.RequestID); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
