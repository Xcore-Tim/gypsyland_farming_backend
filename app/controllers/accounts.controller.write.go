package controllers

import (
	"gypsyland_farming/app/models"

	"net/http"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ctrl AccountRequestController) CreateAccountRequest(ctx *gin.Context) {

	var requestBody models.PostRequestBody

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var accountRequestTask models.AccountRequestTask

	accountRequestTask.AccountRequest = requestBody.AccountRequest

	accountRequestTask.Buyer.ID = requestBody.UserIdentity.UserID
	accountRequestTask.Buyer.Name = requestBody.UserIdentity.FullName
	accountRequestTask.Buyer.Position = requestBody.UserIdentity.RoleID
	accountRequestTask.BuyerID = requestBody.UserIdentity.UserID

	team, err := ctrl.TeamService.GetTeamByNum(requestBody.UserIdentity.TeamID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accountRequestTask.Team.ID = team.ID
	accountRequestTask.Team.Number = team.Number
	accountRequestTask.Team.TeamLead.ID = team.TeamLead.ID
	accountRequestTask.Team.TeamLead.Name = team.TeamLead.Name
	accountRequestTask.Team.TeamLead.Position = team.TeamLead.Position
	accountRequestTask.TeamID = requestBody.UserIdentity.TeamID
	accountRequestTask.Description = requestBody.Description

	if err := ctrl.WriteAccountRequestService.CreateAccountRequest(&accountRequestTask); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (ctrl AccountRequestController) UpdateAccountRequest(ctx *gin.Context) {

	var accountRequestTask models.AccountRequestTask

	if err := ctx.ShouldBindJSON(&accountRequestTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := ctrl.WriteAccountRequestService.UpdateAccountRequest(&accountRequestTask)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": accountRequestTask})
}

func (ctrl AccountRequestController) UpdateRequest(ctx *gin.Context) {

	var accountRequestUpdate models.AccountRequestUpdate

	request_id, _ := primitive.ObjectIDFromHex(ctx.Param("request_id"))

	if err := ctx.ShouldBindJSON(&accountRequestUpdate); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	accountRequestUpdate.ID = request_id

	err := ctrl.WriteAccountRequestService.UpdateRequest(&accountRequestUpdate)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (ctrl AccountRequestController) TakeAccountRequest(ctx *gin.Context) {

	var userIdentity models.UserIdentity

	request_id, _ := primitive.ObjectIDFromHex(ctx.Param("request_id"))

	if err := ctx.ShouldBindJSON(&userIdentity); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	farmer := models.Employee{
		ID:       userIdentity.UserID,
		Name:     userIdentity.FullName,
		Position: userIdentity.RoleID,
	}

	if err := ctrl.WriteAccountRequestService.TakeAccountRequest(&farmer, &request_id); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (ctrl AccountRequestController) CancelAccountRequest(ctx *gin.Context) {

	var canceledRequest models.AccountRequestCanceled

	request_id, _ := primitive.ObjectIDFromHex(ctx.Param("request_id"))

	description := ""

	if err := ctx.ShouldBindJSON(&canceledRequest); err == nil {
		description = canceledRequest.Description
	}

	err := ctrl.WriteAccountRequestService.CancelAccountRequest(&request_id, description)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (ctrl AccountRequestController) CompleteAccountRequest(ctx *gin.Context) {

	var accountRequestCompleted models.AccountRequestCompleted

	request_id, _ := primitive.ObjectIDFromHex(ctx.Param("request_id"))

	if err := ctx.ShouldBindJSON(&accountRequestCompleted); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	accountRequestCompleted.ID = request_id

	err := ctrl.WriteAccountRequestService.CompleteAccountRequest(&accountRequestCompleted)

	if err != nil {
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
