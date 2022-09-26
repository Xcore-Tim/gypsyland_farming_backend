package controllers

import (
	"gypsyland_farming/app/models"
	"math"
	"strconv"
	"strings"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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
	}

	convertPeriod(&requestBody.Period)

	requestBody.Status = status

	requestBody.Convert()

	switch requestBody.UserData.RoleID {

	case 6:
		ctrl.GetFarmerRequests(ctx, &requestBody)
	case 2:

		var accountRequestTasks []models.AccountRequestTask

		if err := ctrl.ReadAccountRequestService.GetTeamleadRequests(&requestBody, &accountRequestTasks); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, accountRequestTasks)
	case 3, 4, 7:
		ctrl.GetBuyerRequests(ctx, &requestBody)
	default:

	}
}

func (ctrl AccountRequestController) GetBuyerRequests(ctx *gin.Context, requestBody *models.GetRequestBody) {

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

func (ctrl AccountRequestController) GetFarmerRequests(ctx *gin.Context, requestBody *models.GetRequestBody) {

	teamAccess, err := ctrl.TeamAccessService.GetAccesses(requestBody.UserData.UserID)

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
			v.Total = roundFloat(v.Total, 2)
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

func (ctrl AccountRequestController) GetTeamleadRequests(ctx *gin.Context, requestBody *models.GetRequestBody) {
	switch requestBody.Status {
	case 0:
		var buyerPendingResponse []models.BuyersPendingResponse
		if err := ctrl.ReadAccountRequestService.GetTeamleadPendingRequests(requestBody, &buyerPendingResponse); err != nil {
			return
		}
		ctx.JSON(http.StatusOK, buyerPendingResponse)
	case 1:
		var buyersInworkReponse []models.BuyersInworkResponse
		if err := ctrl.ReadAccountRequestService.GetTeamleadInworkRequests(requestBody, &buyersInworkReponse); err != nil {
			return
		}
		ctx.JSON(http.StatusOK, buyersInworkReponse)
	case 2:
		var buyersCompletedResponse []models.BuyersCompletedResponse
		if err := ctrl.ReadAccountRequestService.GetTeamleadCompletedRequests(requestBody, &buyersCompletedResponse); err != nil {
			return
		}
		ctx.JSON(http.StatusOK, buyersCompletedResponse)
	case 3:
		var buyersCancelledResponse []models.BuyersCancelledResponse
		if err := ctrl.ReadAccountRequestService.GetTeamleadCancelledRequests(requestBody, &buyersCancelledResponse); err != nil {
			return
		}
		ctx.JSON(http.StatusOK, buyersCancelledResponse)
	}
}

func (ctrl AccountRequestController) AggregateFarmersData(ctx *gin.Context) {

	var groupedResponse []models.GroupedFarmersResponse

	if err := ctrl.ReadAccountRequestService.AggregateFarmersData(&groupedResponse); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var finalResponse []models.GroupedFarmersResponse

	for _, farmer := range groupedResponse {
		teamAccess, err := ctrl.TeamAccessService.GetAccesses(farmer.Farmer.ID)

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
	var groupedResponse []models.GroupedTeamsResponse

	if err := ctrl.ReadAccountRequestService.AggregateTeamsData(&groupedResponse); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, groupedResponse)
}

func (ctrl AccountRequestController) AggregateBuyersData(ctx *gin.Context) {

	teamlead_id_str := ctx.Param("user_id")
	teamlead_id, err := strconv.Atoi(teamlead_id_str)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	results := ctrl.ReadAccountRequestService.AggregateBuyersData(teamlead_id)
	ctx.JSON(http.StatusOK, results)
}

func convertPeriod(period *models.Period) {

	if period.StartISO == "" {
		period.EndDate = time.Now()
	} else {
		date_format := "2006-01-02"
		period.EndDate, _ = time.Parse(date_format, period.EndISO)
		period.StartDate, _ = time.Parse(date_format, period.StartISO)
	}
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
