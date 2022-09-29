package controllers

import (
	"gypsyland_farming/app/models"

	"strconv"
	"strings"

	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ctrl AccountRequestController) GetAll(ctx *gin.Context) {

	accountRequests, err := ctrl.ReadAccountRequestService.GetAll()

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, accountRequests)
}

func (ctrl AccountRequestController) GetPendingRequests(ctx *gin.Context) {

	status := models.Pending
	ctrl.GetAccountRequests(status, ctx)
}

func (ctrl AccountRequestController) GetInworkRequests(ctx *gin.Context) {

	status := models.Inwork
	ctrl.GetAccountRequests(status, ctx)
}

func (ctrl AccountRequestController) GetCompletedRequests(ctx *gin.Context) {

	status := models.Complete
	ctrl.GetAccountRequests(status, ctx)
}

func (ctrl AccountRequestController) GetCanceledRequests(ctx *gin.Context) {

	status := models.Canceled
	ctrl.GetAccountRequests(status, ctx)
}

func (ctrl AccountRequestController) GetAccountRequests(status int, ctx *gin.Context) {

	var requestBody models.GetRequestBody

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	requestBody.Status = status

	requestBody.Convert()

	switch requestBody.UserData.RoleID {

	case 6:
		ctrl.GetFarmerRequests(&requestBody, ctx)
	case 3, 4, 7:
		ctrl.GetBuyerRequests(&requestBody, ctx)
	}
}

func (ctrl AccountRequestController) GetBuyerRequests(requestBody *models.GetRequestBody, ctx *gin.Context) {

	switch requestBody.Status {
	case 0:
		var buyerPendingResponse []models.BuyersPendingResponse
		if err := ctrl.ReadAccountRequestService.GetBuyerPendingRequests(requestBody, &buyerPendingResponse); err != nil {
			return
		}
		ctx.JSON(http.StatusOK, buyerPendingResponse)
	case 1:
		var buyersInworkReponse []models.BuyersInworkResponse
		if err := ctrl.ReadAccountRequestService.GetBuyerInworkRequests(requestBody, &buyersInworkReponse); err != nil {
			return
		}
		ctx.JSON(http.StatusOK, buyersInworkReponse)
	case 2:
		var buyersCompletedResponse []models.BuyersCompletedResponse
		if err := ctrl.ReadAccountRequestService.GetBuyerCompletedRequests(requestBody, &buyersCompletedResponse); err != nil {
			return
		}
		ctx.JSON(http.StatusOK, buyersCompletedResponse)
	case 3:
		var buyersCancelledResponse []models.BuyersCancelledResponse
		if err := ctrl.ReadAccountRequestService.GetBuyerCancelledRequests(requestBody, &buyersCancelledResponse); err != nil {
			return
		}
		ctx.JSON(http.StatusOK, buyersCancelledResponse)
	}
}

func (ctrl AccountRequestController) GetFarmerRequests(requestBody *models.GetRequestBody, ctx *gin.Context) {

	teamAccess, err := ctrl.TeamAccessService.GetAccess(requestBody.UserData.UserID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	switch requestBody.Status {

	case 0:

		var farmersPendingRequest []models.FarmersPendingResponse
		if err := ctrl.ReadAccountRequestService.GetFarmerPeindingRequests(requestBody, &farmersPendingRequest, *teamAccess); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, farmersPendingRequest)

	case 1:
		var farmersInworkRequests []models.FarmersInworkResponse
		if err := ctrl.ReadAccountRequestService.GetFarmerInworkRequests(requestBody, &farmersInworkRequests, *teamAccess); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
		}
		ctx.JSON(http.StatusOK, farmersInworkRequests)

	case 2:
		var farmersCompletedRequests []models.FarmersCompletedResponse
		err := ctrl.ReadAccountRequestService.GetFarmerCompletedRequests(requestBody, &farmersCompletedRequests, *teamAccess)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		for _, v := range farmersCompletedRequests {
			v.Total = ctrl.roundFloat(v.Total, 2)
		}

		ctx.JSON(http.StatusOK, farmersCompletedRequests)
	case 3:
		var farmersCancelledRequests []models.FarmersCancelledResponse
		err := ctrl.ReadAccountRequestService.GetFarmerCancelledRequests(requestBody, &farmersCancelledRequests, *teamAccess)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, farmersCancelledRequests)
	}
}

func (ctrl AccountRequestController) GetAccountRequestData(ctx *gin.Context) {

	requestID, err := primitive.ObjectIDFromHex(ctx.Param("requestID"))

	if err != nil {
		ctx.JSON(http.StatusBadGateway, err.Error())
		return
	}

	accountRequest, err := ctrl.ReadAccountRequestService.GetAccountRequestData(&requestID)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, accountRequest)
}

func (ctrl AccountRequestController) AggregateFarmersData(ctx *gin.Context) {

	var requestBody models.GetRequestBody

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	requestBody.Status = models.Complete

	requestBody.Convert()

	var groupedResponse []models.GroupedFarmersResponse

	if err := ctrl.ReadAccountRequestService.AggregateFarmersData(&requestBody, &groupedResponse); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var finalResponse []models.GroupedFarmersResponse

	for _, farmer := range groupedResponse {
		teamAccess, err := ctrl.TeamAccessService.GetAccess(farmer.Farmer.ID)

		if err != nil {
			continue
		}

		var builder strings.Builder

		for i, v := range teamAccess.Teams {
			value := strconv.Itoa(v)

			if i == 0 {
				builder.WriteString(value)
			} else {
				builder.WriteString(", " + value)
			}

		}

		farmer.Teams = builder.String()
		finalResponse = append(finalResponse, farmer)
	}

	ctx.JSON(http.StatusOK, finalResponse)
}

func (ctrl AccountRequestController) AggregateTeamsData(ctx *gin.Context) {

	var requestBody models.GetRequestBody

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	requestBody.Status = models.Complete

	requestBody.Convert()

	var groupedResponse []models.GroupedTeamsResponse

	if err := ctrl.ReadAccountRequestService.AggregateTeamsData(&requestBody, &groupedResponse); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, groupedResponse)
}

func (ctrl AccountRequestController) AggregateBuyersData(ctx *gin.Context) {

	var requestBody models.GetRequestBody

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	requestBody.Status = models.Complete

	requestBody.Convert()

	var response []models.GroupedBuyersResponse

	if err := ctrl.ReadAccountRequestService.AggregateBuyersData(&requestBody, &response, requestBody.UserData.UserID); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, response)
}
