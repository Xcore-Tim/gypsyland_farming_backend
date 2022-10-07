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

func (ctrl TeamAccessController) GetAllAccesses(ctx *gin.Context) {

	teamAccesses, err := ctrl.TeamAccessService.GetAllAccesses()

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, teamAccesses)
}

func (ctrl TeamAccessController) GetFarmersAccesses(ctx *gin.Context) {

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

	user_id_str := ctx.Param("user_id")
	user_id, err := strconv.Atoi(string(user_id_str))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	teamAccess, err := ctrl.TeamAccessService.GetAccess(user_id)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, teamAccess)
}

func (ctrl TeamAccessController) RegisterUserRoutes(rg *gin.RouterGroup) {

	teamsGroup := rg.Group("/team")

	accessGroup := teamsGroup.Group("/access")
	accessGroup.POST("/getall", ctrl.GetAllAccesses)
	accessGroup.POST("/farmers", ctrl.GetFarmersAccesses)

	accessGroup.POST("/get/:user_id", ctrl.GetAccesses)
	accessGroup.POST("/add", ctrl.AddAccess)
	accessGroup.POST("/revoke", ctrl.RevokeAccess)
}
