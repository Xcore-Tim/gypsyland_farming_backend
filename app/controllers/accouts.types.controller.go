package controllers

import (
	"github.com/gin-gonic/gin"

	"gypsyland_farming/app/models"
	"gypsyland_farming/app/services"
	"net/http"
)

type AccountTypesController struct {
	AccountTypesService services.AccountTypesService
}

func NewAccountTypesController(

	accountTypesService services.AccountTypesService) AccountTypesController {

	return AccountTypesController{
		AccountTypesService: accountTypesService,
	}
}

func (atc AccountTypesController) CreateAccountType(ctx *gin.Context) {

	var accountType models.AccountType

	if err := ctx.ShouldBindJSON(&accountType); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := atc.AccountTypesService.CreateAccountType(&accountType)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": accountType})
}

func (atc AccountTypesController) GetAllTypes(ctx *gin.Context) {

	accountTypes, err := atc.AccountTypesService.GetAllAccountTypes()

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, accountTypes)
}

func (atc AccountTypesController) RegisterUserRoutes(rg *gin.RouterGroup) {

	accountTypeGroup := rg.Group("/accountTypes")

	accountTypeGroup.POST("/create", atc.CreateAccountType)
	accountTypeGroup.POST("/getall", atc.GetAllTypes)
}
