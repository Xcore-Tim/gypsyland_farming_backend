package controllers

import (
	teams "gypsylandFarming/app/models/teams"
	teamService "gypsylandFarming/app/services/teams"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TeamController struct {
	TeamService       teamService.TeamService
	TeamAccessService teamService.TeamAccessService
}

func NewTeamController(teamService teamService.TeamService, teamAccessService teamService.TeamAccessService) TeamController {
	return TeamController{
		TeamService:       teamService,
		TeamAccessService: teamAccessService,
	}
}

func (ctrl TeamController) CreateTeam(ctx *gin.Context) {

	var team teams.Team

	if err := ctx.ShouldBindJSON(&team); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	err := ctrl.TeamService.CreateTeam(&team)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (ctrl TeamController) GetAllTeams(ctx *gin.Context) {

	teams, err := ctrl.TeamService.GetAllTeams()

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, teams)

}

func (ctrl TeamController) GetDropdown(ctx *gin.Context) {

	var editAccessRequest teams.EditTeamAccessRequest

	if err := ctx.ShouldBindJSON(&editAccessRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	editAccessRequest.Convert()

	var teamAccess teams.TeamAccess
	teamAccess.Teams = append(teamAccess.Teams, 0)

	ctrl.TeamAccessService.GetAccessByNum(&teamAccess, editAccessRequest.UserData.UserID)

	teams, err := ctrl.TeamService.GetDropdown(&teamAccess, &editAccessRequest)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, teams)
}

func (ctrl TeamController) ImportTeams(ctx *gin.Context) {

	token := ctx.Request.PostFormValue("token")

	if err := ctrl.TeamService.ImportTeams(token); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (ctrl TeamController) RegisterUserRoutes(rg *gin.RouterGroup) {

	teamsGroup := rg.Group("/team")

	teamsGroup.POST("/create", ctrl.CreateTeam)
	teamsGroup.POST("/import", ctrl.ImportTeams)

	getGroup := teamsGroup.Group("/get")
	getGroup.POST("/dropdown", ctrl.GetDropdown)
	getGroup.POST("/all", ctrl.GetAllTeams)
}
