package controllers

import (
	"gypsyland_farming/app/models"
	"gypsyland_farming/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TeamController struct {
	TeamService       services.TeamService
	TeamAccessService services.TeamAccessService
	EmployeeService   services.EmployeeService
}

func NewTeamController(teamService services.TeamService) TeamController {
	return TeamController{
		TeamService: teamService,
	}
}

func (ctrl TeamController) CreateTeam(ctx *gin.Context) {

	var team models.Team

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

func (ctrl TeamController) ImportTeams(ctx *gin.Context) {

	token := ctx.Request.PostFormValue("token")

	if err := ctrl.TeamService.ImportTeams(token); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})

	// resp := ctrl.TeamService.ImportTeams1(token)
	// ctx.JSON(http.StatusAccepted, resp)

}

func (ctrl TeamController) RegisterUserRoutes(rg *gin.RouterGroup) {

	teamsGroup := rg.Group("/team")

	teamsGroup.POST("/create", ctrl.CreateTeam)
	teamsGroup.POST("/getall", ctrl.GetAllTeams)
	teamsGroup.POST("/import", ctrl.ImportTeams)

}
