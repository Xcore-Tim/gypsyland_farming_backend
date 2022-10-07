package controllers

import (
	teams "gypsylandFarming/app/models/teams"
	teamService "gypsylandFarming/app/services/teams"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TeamAccessController struct {
	TeamAccessService teamService.TeamAccessService
}

func NewTeamAccessController(teamAccessService teamService.TeamAccessService) TeamAccessController {
	return TeamAccessController{
		TeamAccessService: teamAccessService,
	}
}

func (ctrl TeamAccessController) AddAccess(ctx *gin.Context) {

	var editAccessRequest teams.EditTeamAccessRequest

	if err := ctx.ShouldBindJSON(&editAccessRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	editAccessRequest.Convert()

	if err := ctrl.TeamAccessService.AddAccess(editAccessRequest.UserData.UserID, editAccessRequest.TeamEdit.TeamID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (ctrl TeamAccessController) RevokeAccess(ctx *gin.Context) {

	var editAccessRequest teams.EditTeamAccessRequest

	if err := ctx.ShouldBindJSON(&editAccessRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	editAccessRequest.Convert()

	if err := ctrl.TeamAccessService.RevokeAccess(editAccessRequest.UserData.UserID, editAccessRequest.TeamEdit.TeamID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (ctrl TeamAccessController) GetFarmerAccesses(ctx *gin.Context) {

	var farmerAccessesRequest teams.EditTeamAccessRequest

	if err := ctx.ShouldBindJSON(&farmerAccessesRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	farmerAccessesRequest.Convert()

	var farmerAccesses []teams.FarmerAccess

	if err := ctrl.TeamAccessService.GetFarmersAccesses(&farmerAccesses, &farmerAccessesRequest.UserData); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, farmerAccesses)
}

func (ctrl TeamAccessController) GetAccesses(ctx *gin.Context) {

	strUserID := ctx.Param("userID")
	userID, err := strconv.Atoi(string(strUserID))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	teamAccess, err := ctrl.TeamAccessService.GetAccess(userID)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, teamAccess)
}

func (ctrl TeamAccessController) GetAll(ctx *gin.Context) {

	teamAccesses, err := ctrl.TeamAccessService.GetAllAccesses()

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, teamAccesses)
}

func (ctrl TeamAccessController) RegisterUserRoutes(rg *gin.RouterGroup) {

	teamsGroup := rg.Group("/team")

	accessGroup := teamsGroup.Group("/access")

	accessGroup.POST("/farmers", ctrl.GetFarmerAccesses)

	accessGroup.POST("/add", ctrl.AddAccess)
	accessGroup.POST("/revoke", ctrl.RevokeAccess)

	getGroup := accessGroup.Group("/get")
	getGroup.POST("/all", ctrl.GetAll)
	getGroup.POST("/access", ctrl.GetAccesses)

}
