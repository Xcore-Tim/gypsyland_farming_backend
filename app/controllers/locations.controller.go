package controllers

import (
	"gypsyland_farming/app/models"
	"gypsyland_farming/app/services"
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

func (lc *LocationController) CreateLocation(ctx *gin.Context) {

	var location models.Location
	if err := ctx.ShouldBindJSON(&location); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := lc.LocationService.CreateLocation(&location)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (lc *LocationController) GetLocation(ctx *gin.Context) {

	locationId, _ := primitive.ObjectIDFromHex(ctx.Param("id"))
	location, err := lc.LocationService.GetLocation(&locationId)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, location)
}

func (lc LocationController) GetAll(ctx *gin.Context) {

	users, err := lc.LocationService.GetAll()

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (lc LocationController) UpdateLocation(ctx *gin.Context) {

	var location models.Location

	if err := ctx.ShouldBindJSON(&location); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := lc.LocationService.UpdateLocation(&location)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, location)
}

func (lc LocationController) DeleteLocation(ctx *gin.Context) {

	name := ctx.Param("name")

	err := lc.LocationService.DeleteLocation(&name)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (lc LocationController) RegisterUserRoutes(rg *gin.RouterGroup) {

	locationRoute := rg.Group("/locations")
	locationRoute.POST("/create", lc.CreateLocation)
	locationRoute.POST("/get/:id", lc.GetLocation)
	locationRoute.POST("/getall", lc.GetAll)
	locationRoute.POST("/update", lc.UpdateLocation)
	locationRoute.POST("/delete/:name", lc.DeleteLocation)

}
