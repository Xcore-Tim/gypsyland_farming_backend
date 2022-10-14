package controllers

import (
	accountTypesService "gypsylandFarming/app/services/accountTypes"
	accountService "gypsylandFarming/app/services/accounts"
	currencyServices "gypsylandFarming/app/services/currency"
	fileService "gypsylandFarming/app/services/files"
	services "gypsylandFarming/app/services/locations"
	teamService "gypsylandFarming/app/services/teams"

	"github.com/gin-gonic/gin"
)

type AccountRequestController struct {
	ReadAccountRequestService  accountService.ReadAccountRequestService
	WriteAccountRequestService accountService.WriteAccountRequestService
	AccountTypesService        accountTypesService.AccountTypesService
	TeamAccessService          teamService.TeamAccessService
	TeamService                teamService.TeamService
	LocationService            services.LocationService
	FileService                fileService.FileService
	CurrencyService            currencyServices.CurrencyService
	CurrencyRatesService       currencyServices.CurrencyRatesService
}

type UserDataConverter interface {
	Convert()
}

func NewAccountRequestTaskController(
	readAccountRequestService accountService.ReadAccountRequestService,
	writeAccountRequestService accountService.WriteAccountRequestService,
	teamService teamService.TeamService,
	teamAccessService teamService.TeamAccessService,
	locationService services.LocationService,
	accountTypesService accountTypesService.AccountTypesService,
	fileService fileService.FileService,
	currencyService currencyServices.CurrencyService,
	currencyRatesService currencyServices.CurrencyRatesService,

) AccountRequestController {
	return AccountRequestController{
		ReadAccountRequestService:  readAccountRequestService,
		WriteAccountRequestService: writeAccountRequestService,
		TeamService:                teamService,
		TeamAccessService:          teamAccessService,
		LocationService:            locationService,
		AccountTypesService:        accountTypesService,
		FileService:                fileService,
		CurrencyService:            currencyService,
		CurrencyRatesService:       currencyRatesService,
	}
}

func (ctrl AccountRequestController) RegisterUserRoutes(rg *gin.RouterGroup) {

	accountRequestGroup := rg.Group("/accountRequests")

	accountRequestGroup.POST("/create", ctrl.CreateAccountRequest)

	getGroup := accountRequestGroup.Group("/get")
	getGroup.GET("/request", ctrl.GetAccountRequestData)
	getGroup.POST("/all", ctrl.GetAll)

	getGroup.POST("/inwork", ctrl.GetInworkRequests)
	getGroup.POST("/pending", ctrl.GetPendingRequests)
	getGroup.POST("/completed", ctrl.GetCompletedRequests)
	getGroup.POST("/canceled", ctrl.GetCanceledRequests)

	aggregatedGroup := accountRequestGroup.Group("/aggregate")
	aggregatedGroup.POST("/farmers", ctrl.AggregateFarmersData)
	aggregatedGroup.POST("/teams", ctrl.AggregateTeamsData)
	aggregatedGroup.POST("/buyers", ctrl.AggregateBuyersData)

	updateGroup := accountRequestGroup.Group("/update")
	updateGroup.POST("/request", ctrl.UpdateRequest)

	statusGroup := updateGroup.Group("/status")
	statusGroup.POST("/inwork", ctrl.TakeAccountRequest)
	statusGroup.POST("/canceled", ctrl.CancelAccountRequest)
	statusGroup.POST("/completed", ctrl.CompleteAccountRequest)
	statusGroup.POST("/return", ctrl.ReturnAccountRequest)

	deleteGroup := accountRequestGroup.Group("/delete")
	deleteGroup.POST("/all", ctrl.DeleteAllAccountRequests)
	deleteGroup.POST("/request", ctrl.DeleteAccountRequest)

	accountRequestGroup.POST("/test", ctrl.Test)
}
