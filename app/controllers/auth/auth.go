package controllers

import (
	"errors"
	auth "gypsylandFarming/app/models/authentication"

	authenticationService "gypsylandFarming/app/services/auth"
	teamService "gypsylandFarming/app/services/teams"

	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	JWTService  authenticationService.JWTService
	AuthService authenticationService.AuthService
	TeamService teamService.TeamService
}

func NewAuthController(jwtService authenticationService.JWTService, authService authenticationService.AuthService, teamService teamService.TeamService) AuthController {
	return AuthController{
		JWTService:  jwtService,
		AuthService: authService,
		TeamService: teamService,
	}
}

func (ctrl AuthController) Login(ctx *gin.Context) {

	var authData auth.UserCredentials

	if err := ctx.ShouldBindJSON(&authData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var authResponse auth.AuthResponseData

	if err := ctrl.AuthService.Login(&authData, &authResponse); err != nil {
		badCredentials := errors.New("incorrect user credentials")
		ctrl.AuthService.AuthError(&authResponse, badCredentials.Error())
		ctx.JSON(http.StatusBadRequest, authResponse)
		return
	}

	ctx.JSON(http.StatusOK, authResponse)
}

func (ctrl AuthController) ValidateToken(ctx *gin.Context) {

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiIxMDUiLCJSb2xlSWQiOiI2IiwiVGVhbUlkIjoiIiwibmJmIjoxNjY0NDYwODI0LCJleHAiOjE2NjQ0ODk2MjQsImlzcyI6Ik15QXV0aFNlcnZlciIsImF1ZCI6Ik15QXV0aENsaWVudCJ9.czNbSDWdWEAv3N4e0mfGps6EwpLx2risUSvGz5JUuqs"
	res, _ := ctrl.JWTService.ValidateToken(token)
	ctx.JSON(http.StatusAccepted, res)
}

func (ctrl AuthController) RegisterUserRoutes(rg *gin.RouterGroup) {

	authRequestGroup := rg.Group("/auth")

	authRequestGroup.POST("", ctrl.Login)
	authRequestGroup.POST("/validate", ctrl.ValidateToken)
}
