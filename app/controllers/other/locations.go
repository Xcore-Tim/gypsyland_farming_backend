package controllers

import (
	"gypsylandFarming/app/models"
	services "gypsylandFarming/app/services/other"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LocationController struct {
	LocationService services.LocationService
}

func NewLocationController(locationService services.LocationService) LocationController {
	return LocationController{
		LocationService: locationService,
	}
}

func (ctrl *LocationController) CreateLocation(ctx *gin.Context) {

	var location models.Location

	if err := ctx.ShouldBindJSON(&location); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := ctrl.LocationService.CreateLocation(&location)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (ctrl *LocationController) GetLocation(ctx *gin.Context) {

	locationId, _ := primitive.ObjectIDFromHex(ctx.Query("id"))
	location, err := ctrl.LocationService.GetLocation(locationId)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, location)
}

func (ctrl LocationController) GetAll(ctx *gin.Context) {

	users, err := ctrl.LocationService.GetAll()

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (ctrl LocationController) UpdateLocation(ctx *gin.Context) {

	oid, err := primitive.ObjectIDFromHex(ctx.Query("oid"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	var location models.Location

	if err := ctx.ShouldBindJSON(&location); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err = ctrl.LocationService.UpdateLocation(&oid, &location); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, location)
}

func (ctrl LocationController) DeleteLocation(ctx *gin.Context) {

	oid := ctx.Query("oid")

	err := ctrl.LocationService.DeleteLocation(&oid)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (ctrl LocationController) DeleteAll(ctx *gin.Context) {

	typesCount, err := ctrl.LocationService.DeleteAll()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"deleted": typesCount})
}

func (ctrl LocationController) RegisterUserRoutes(rg *gin.RouterGroup) {

	locationRoute := rg.Group("/locations")
	locationRoute.POST("/create", ctrl.CreateLocation)

	getGroup := locationRoute.Group("/get")
	getGroup.POST("/location", ctrl.GetLocation)
	getGroup.POST("/all", ctrl.GetAll)

	locationRoute.POST("/update", ctrl.UpdateLocation)

	deleteGroup := locationRoute.Group("delete")
	deleteGroup.POST("/location", ctrl.DeleteLocation)
	deleteGroup.POST("/all")

}
