package controllers

import (
	"gypsyland_farming/app/models"
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

func (ctrl AccountRequestController) GetAccountRequests(status int, ctx *gin.Context) {

	var requestBody models.GetRequestBody

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	convertPeriod(&requestBody.Period)

	requestBody.Status = status

	switch requestBody.UserIdentity.RoleID {

	case 6:
		teamAccess, err := ctrl.TeamAccessService.GetAccesses(requestBody.UserIdentity.UserID)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "found no access for farmer"})
			return
		}

		var accountRequestTasks []models.AccountRequestTask

		err = ctrl.ReadAccountRequestService.GetFarmerRequests(&requestBody, &accountRequestTasks, teamAccess)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "found no documents"})
			return
		}

		ctx.JSON(http.StatusOK, accountRequestTasks)

	case 2:

		var accountRequestTasks []models.AccountRequestTask

		err := ctrl.ReadAccountRequestService.GetTeamleadRequests(&requestBody, &accountRequestTasks)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, accountRequestTasks)

	case 1, 5:

		var accountRequestTasks []models.AccountRequestTask

		err := ctrl.ReadAccountRequestService.GetTLFARequests(&requestBody, &accountRequestTasks)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "found no documents"})
			return
		}

		ctx.JSON(http.StatusOK, accountRequestTasks)

	}
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
