package controllers

import (
	"gypsyland_farming/app/models"

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

	var requestBody models.CreateAccountRequestBody

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}

	requestBody.Convert()

	var accountRequestTask models.AccountRequestTask

	location, err := ctrl.LocationService.GetLocation(requestBody.AccountRequestData.LocationID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
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

	accountRequestTask.AccountRequest.Location = *location
	accountRequestTask.AccountRequest.Type = *accountType
	accountRequestTask.AccountRequest.Quantity = requestBody.AccountRequestData.Quantity

	accountRequestTask.DateCreated = time.Now().Unix()
	accountRequestTask.Description = requestBody.AccountRequestBody.Description
	accountRequestTask.Price = requestBody.AccountRequestData.Price

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

	var accountRequestUpdate models.UpdateRequestBody

	if err := ctx.ShouldBindJSON(&accountRequestUpdate); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	accountRequestUpdate.Convert()

	accountRequest, err := ctrl.ReadAccountRequestService.GetAccountRequestData(&accountRequestUpdate.RequestID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	accountRequestUpdate.AccountRequest = *accountRequest

	if accountRequestUpdate.UpdateBody.Location != "" {
		if location, err := ctrl.LocationService.GetLocationByName(accountRequestUpdate.UpdateBody.Location); err == nil {
			accountRequestUpdate.AccountRequest.Location = *location
		}
	}

	if accountRequestUpdate.UpdateBody.AccountType != "" {
		if accountType, err := ctrl.AccountTypesService.GetTypeByName(accountRequestUpdate.UpdateBody.AccountType); err == nil {
			accountRequestUpdate.AccountRequest.Type = *accountType
		}
	}

	if accountRequestUpdate.UpdateBody.Quantity != accountRequest.Quantity {
		accountRequestUpdate.AccountRequest.Quantity = accountRequestUpdate.UpdateBody.Quantity
	}

	ctx.JSON(http.StatusOK, accountRequestUpdate)

	if err := ctrl.WriteAccountRequestService.UpdateRequest(&accountRequestUpdate); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (ctrl AccountRequestController) TakeAccountRequest(ctx *gin.Context) {

	var requestData models.TakeAccountRequest

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

	var cancelRequest models.CancelAccountRequest

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

	var accountRequestCompleted models.CompleteAccountRequest

	if err := ctx.ShouldBindJSON(&accountRequestCompleted); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	accountRequestCompleted.Convert()

	if accountRequest, err := ctrl.ReadAccountRequestService.GetRequest(&accountRequestCompleted.RequestID); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	} else {
		total := float64(accountRequest.AccountRequest.Quantity) * accountRequestCompleted.OrderInfo.Price
		accountRequestCompleted.TotalSum = ctrl.roundFloat(total, 2)
	}

	if err := ctrl.WriteAccountRequestService.CompleteAccountRequest(&accountRequestCompleted); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (ctrl AccountRequestController) ReturnAccountRequest(ctx *gin.Context) {

	requestID, _ := primitive.ObjectIDFromHex(ctx.Param("requestID"))

	if err := ctrl.WriteAccountRequestService.ReturnAccountRequest(&requestID); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (ctrl AccountRequestController) DeleteAccountRequest(ctx *gin.Context) {

	request_id, err := primitive.ObjectIDFromHex(ctx.Param("request_id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = ctrl.WriteAccountRequestService.DeleteAccountRequest(&request_id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
