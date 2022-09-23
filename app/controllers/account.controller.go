package controllers

import (
	"gypsyland_farming/app/services"

	"github.com/gin-gonic/gin"
)

type AccountRequestController struct {
	ReadAccountRequestService  services.ReadAccountRequestService
	WriteAccountRequestService services.WriteAccountRequestService
	TeamAccessService          services.TeamAccessService
	TeamService                services.TeamService
	LocationService            services.LocationService
	AccountTypesService        services.AccountTypesService
}

func NewAccountRequestTaskController(
	readAccountRequestService services.ReadAccountRequestService,
	writeAccountRequestService services.WriteAccountRequestService,
	teamService services.TeamService,
	teamAccessService services.TeamAccessService,
	locationService services.LocationService,
	accountTypesService services.AccountTypesService,
) AccountRequestController {
	return AccountRequestController{
		ReadAccountRequestService:  readAccountRequestService,
		WriteAccountRequestService: writeAccountRequestService,
		TeamService:                teamService,
		TeamAccessService:          teamAccessService,
		LocationService:            locationService,
		AccountTypesService:        accountTypesService,
	}
}

func (ctrl AccountRequestController) RegisterUserRoutes(rg *gin.RouterGroup) {

	accountRequestGroup := rg.Group("/accountRequests")

	accountRequestGroup.POST("/getall", ctrl.GetAll)

	getGroup := accountRequestGroup.Group("/get")

	getGroup.POST("/inwork", ctrl.GetInworkRequests)
	getGroup.POST("/pending", ctrl.GetPendingRequests)
	getGroup.POST("/completed", ctrl.GetCompletedRequests)
	getGroup.POST("/canceled", ctrl.GetCanceledRequests)

	aggregatedGroup := accountRequestGroup.Group("/aggregate")

	aggregatedGroup.POST("/farmers", ctrl.AggregateFarmersData)
	aggregatedGroup.POST("/teams", ctrl.AggregateTeamsData)
	aggregatedGroup.POST("/buyers/:user_id", ctrl.AggregateBuyersData)

	accountRequestGroup.POST("/create", ctrl.CreateAccountRequest)

	updateGroup := accountRequestGroup.Group("/update")

	updateGroup.POST("/request", ctrl.UpdateRequest)

	statusGroup := updateGroup.Group("/status")

	statusGroup.POST("/inwork", ctrl.TakeAccountRequest)
	statusGroup.POST("/canceled", ctrl.CancelAccountRequest)
	statusGroup.POST("/completed/:request_id", ctrl.CompleteAccountRequest)
	statusGroup.POST("/return/:request_id/:user_id", ctrl.ReturnAccountRequest)

	deleteGroup := accountRequestGroup.Group("/delete")
	deleteGroup.POST("/:request_id", ctrl.DeleteAccountRequest)
}

type UserDataConverter interface {
	Convert()
}
