package controllers

import (
	"gypsyland_farming/app/models"
	"gypsyland_farming/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PositionController struct {
	PositionService services.PositionService
}

func NewPositionController(positionService services.PositionService) PositionController {
	return PositionController{
		PositionService: positionService,
	}
}

func (pc *PositionController) CreatePosition(ctx *gin.Context) {

	var position models.Position

	if err := ctx.ShouldBindJSON(&position); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := pc.PositionService.CreatePosition(&position)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})

}

func (pc *PositionController) GetPosition(ctx *gin.Context) {

	positionId, _ := primitive.ObjectIDFromHex(ctx.Param("_id"))
	position, err := pc.PositionService.GetPosition(&positionId)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": position})
}

func (pc *PositionController) GetAll(ctx *gin.Context) {

	positions, err := pc.PositionService.GetAll()

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}

	ctx.JSON(http.StatusOK, positions)
}

func (pc PositionController) UpdatePosition(ctx *gin.Context) {

	var position models.Position

	if err := ctx.ShouldBindJSON(&position); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := pc.PositionService.UpdatePosition(&position)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": position})
}

func (pc *PositionController) DeletePosition(ctx *gin.Context) {

	id, _ := primitive.ObjectIDFromHex(ctx.Param("_id"))
	err := pc.PositionService.DeletePosition(&id)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (pc PositionController) RegisterUserRoutes(rg *gin.RouterGroup) {

	positionRoute := rg.Group("/positions")
	positionRoute.POST("/create", pc.CreatePosition)
	positionRoute.POST("/get/:_id", pc.GetPosition)
	positionRoute.POST("/getall", pc.GetAll)
	positionRoute.POST("/update", pc.UpdatePosition)
	positionRoute.POST("/delete/:_id", pc.DeletePosition)

}
