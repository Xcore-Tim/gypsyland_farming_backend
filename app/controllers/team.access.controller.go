package controllers

import (
	"gypsyland_farming/app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TeamAccessController struct {
	TeamAccessService services.TeamAccessService
}

func NewTeamAccessController(teamAccessService services.TeamAccessService) TeamAccessController {
	return TeamAccessController{
		TeamAccessService: teamAccessService,
	}
}

func (ctrl TeamAccessController) AddAccess(ctx *gin.Context) {

	team_id_str := ctx.Param("team_number")
	team_id, err := strconv.Atoi(team_id_str)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user_id_str := ctx.Param("user_id")
	user_id, err := strconv.Atoi(user_id_str)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = ctrl.TeamAccessService.AddAccess(user_id, team_id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (ctrl TeamAccessController) RevokeAccess(ctx *gin.Context) {

	team_id_str := ctx.Param("team_number")
	team_id, err := strconv.Atoi(team_id_str)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user_id_str := ctx.Param("user_id")
	user_id, err := strconv.Atoi(user_id_str)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = ctrl.TeamAccessService.RevokeAccess(user_id, team_id)

	if err != nil {
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

func (ctrl TeamAccessController) GetAccesses(ctx *gin.Context) {

	user_id_str := ctx.Param("user_id")
	user_id, err := strconv.Atoi(string(user_id_str))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	teamAccess, err := ctrl.TeamAccessService.GetAccesses(user_id)

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
	accessGroup.POST("/get/:user_id", ctrl.GetAccesses)
	accessGroup.POST("/add/:user_id/:team_number", ctrl.AddAccess)
	accessGroup.POST("/revoke/:user_id/:team_number", ctrl.RevokeAccess)

}
