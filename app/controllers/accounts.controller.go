package controllers

import (
	"gypsyland_farming/app/services"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type AccountRequestController struct {
	ReadAccountRequestService  services.ReadAccountRequestService
	WriteAccountRequestService services.WriteAccountRequestService
	TeamAccessService          services.TeamAccessService
	TeamService                services.TeamService
	LocationService            services.LocationService
	AccountTypesService        services.AccountTypesService
	FileService                services.FileService
}

type UserDataConverter interface {
	Convert()
}

func NewAccountRequestTaskController(
	readAccountRequestService services.ReadAccountRequestService,
	writeAccountRequestService services.WriteAccountRequestService,
	teamService services.TeamService,
	teamAccessService services.TeamAccessService,
	locationService services.LocationService,
	accountTypesService services.AccountTypesService,
	fileService services.FileService,
) AccountRequestController {
	return AccountRequestController{
		ReadAccountRequestService:  readAccountRequestService,
		WriteAccountRequestService: writeAccountRequestService,
		TeamService:                teamService,
		TeamAccessService:          teamAccessService,
		LocationService:            locationService,
		AccountTypesService:        accountTypesService,
		FileService:                fileService,
	}
}

func (ctrl AccountRequestController) RegisterUserRoutes(rg *gin.RouterGroup) {

	accountRequestGroup := rg.Group("/accountRequests")

	getGroup := accountRequestGroup.Group("/get")
	getGroup.GET("/request/:requestID", ctrl.GetAccountRequestData)
	getGroup.POST("/all", ctrl.GetAll)

	getGroup.POST("/inwork", ctrl.GetInworkRequests)
	getGroup.POST("/pending", ctrl.GetPendingRequests)
	getGroup.POST("/completed", ctrl.GetCompletedRequests)
	getGroup.POST("/canceled", ctrl.GetCanceledRequests)

	aggregatedGroup := accountRequestGroup.Group("/aggregate")
	aggregatedGroup.POST("/farmers", ctrl.AggregateFarmersData)
	aggregatedGroup.POST("/teams", ctrl.AggregateTeamsData)
	aggregatedGroup.POST("/buyers", ctrl.AggregateBuyersData)

	accountRequestGroup.POST("/create", ctrl.CreateAccountRequest)

	updateGroup := accountRequestGroup.Group("/update")
	updateGroup.POST("/request", ctrl.UpdateRequest)

	statusGroup := updateGroup.Group("/status")
	statusGroup.POST("/inwork", ctrl.TakeAccountRequest)
	statusGroup.POST("/canceled", ctrl.CancelAccountRequest)
	statusGroup.POST("/completed", ctrl.CompleteAccountRequest)
	statusGroup.POST("/return", ctrl.ReturnAccountRequest)

	deleteGroup := accountRequestGroup.Group("/delete")
	deleteGroup.POST("/all", ctrl.DeleteAllAccountRequests)

	accountRequestGroup.POST("/test", ctrl.Test)
}

func (ctrl AccountRequestController) DeleteAllAccountRequests(ctx *gin.Context) {

	requestCount, err := ctrl.WriteAccountRequestService.DeleteAccountRequests()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	tasksCount, err := ctrl.WriteAccountRequestService.DeleteAccountRequestTasks()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	path := "/var/www/html/react/downloads"

	dir, err := os.ReadDir(path)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusConflict, err.Error())
	}

	for _, d := range dir {
		os.RemoveAll(filepath.Join([]string{"tmp", d.Name()}...))
	}

	ctx.JSON(http.StatusOK, gin.H{"account requests": requestCount, "request tasks": tasksCount})
}
