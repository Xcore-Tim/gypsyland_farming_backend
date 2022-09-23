package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

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

func (ctrl AccountTypesController) CreateAccountType(ctx *gin.Context) {

	var accountType models.AccountType

	if err := ctx.ShouldBindJSON(&accountType); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := ctrl.AccountTypesService.CreateAccountType(&accountType)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": accountType})
}

func (ctrl AccountTypesController) GetAll(ctx *gin.Context) {

	accountTypes, err := ctrl.AccountTypesService.GetAll()

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, accountTypes)
}

func (ctrl AccountTypesController) GetType(ctx *gin.Context) {

	accountTypeID, _ := primitive.ObjectIDFromHex(ctx.Param("accountTypeID"))

	if accountType, err := ctrl.AccountTypesService.GetType(accountTypeID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		ctx.JSON(http.StatusOK, accountType)
	}

}

func (ctrl AccountTypesController) RegisterUserRoutes(rg *gin.RouterGroup) {

	accountTypeGroup := rg.Group("/accountTypes")

	accountTypeGroup.POST("/create", ctrl.CreateAccountType)
	accountTypeGroup.POST("/getall", ctrl.GetAll)
	accountTypeGroup.POST("/get/:accountTypeID", ctrl.GetType)
}
