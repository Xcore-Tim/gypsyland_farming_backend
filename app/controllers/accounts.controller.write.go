package controllers

import (
	"gypsyland_farming/app/models"
	"strconv"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ctrl AccountRequestController) CreateAccountRequest(ctx *gin.Context) {

	var requestBody models.CreateAccountRequestBody

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var accountRequestTask models.AccountRequestTask

	if locationID, err := primitive.ObjectIDFromHex(requestBody.AccountRequest.LocationID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": locationID})
		return
	} else {
		if location, err := ctrl.LocationService.GetLocation(locationID); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": location})
			return
		} else {
			accountRequestTask.AccountRequest.Location = *location
		}
	}

	if accountTypeID, err := primitive.ObjectIDFromHex(requestBody.AccountRequest.TypeID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": accountTypeID})
		return
	} else {
		if accountType, err := ctrl.AccountTypesService.GetType(accountTypeID); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": accountType})
			return
		} else {
			accountRequestTask.AccountRequest.Type = *accountType
		}
	}

	accountRequestTask.AccountRequest.Quantity, _ = strconv.Atoi(requestBody.AccountRequest.Quantity)

	requestBody.Convert()

	team, err := ctrl.TeamService.GetTeamByNum(requestBody.UserData.TeamID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accountRequestTask.Team = *team
	accountRequestTask.DateCreated = time.Now().Unix()
	accountRequestTask.Description = requestBody.AccountRequest.Description

	accountRequestTask.Buyer.ID = requestBody.UserData.UserID
	accountRequestTask.Buyer.Name = requestBody.UserData.Username
	accountRequestTask.Buyer.Position = requestBody.UserData.RoleID

	if err := ctrl.WriteAccountRequestService.CreateAccountRequest(&accountRequestTask); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (ctrl AccountRequestController) UpdateRequest(ctx *gin.Context) {

	var accountRequestUpdate models.UpdateAccountRequest

	if err := ctx.ShouldBindJSON(&accountRequestUpdate); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	accountRequestUpdate.Convert()

	if err := ctrl.WriteAccountRequestService.UpdateRequestNew(&accountRequestUpdate); err != nil {
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
		accountRequestCompleted.TotalSum = float64(accountRequest.AccountRequest.Quantity) * accountRequestCompleted.OrderInfo.Price
	}

	if err := ctrl.WriteAccountRequestService.CompleteAccountRequest(&accountRequestCompleted); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (ctrl AccountRequestController) ReturnAccountRequest(ctx *gin.Context) {

	request_id, _ := primitive.ObjectIDFromHex(ctx.Param("request_id"))

	accountRequestTask, err := ctrl.WriteAccountRequestService.ReturnAccountRequest(&request_id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	accountRequestTask.AccountRequest.ID = primitive.NewObjectID()

	err = ctrl.WriteAccountRequestService.CreateAccountRequest(accountRequestTask)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": accountRequestTask.AccountRequest.ID})

	err = ctrl.WriteAccountRequestService.DeleteAccountRequest(&request_id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"deleted": err.Error()})
		return
	}
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
