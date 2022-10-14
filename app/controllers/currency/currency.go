package controllers

import (
	"errors"
	currencyModels "gypsylandFarming/app/models/currency"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ctrl CurrencyController) CreateCurrency(ctx *gin.Context) {

	var currency currencyModels.Currency

	if err := ctx.ShouldBindJSON(&currency); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	oid, err := ctrl.CurrencyService.CreateCurrency(&currency)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": oid})
}

func (ctrl CurrencyController) GetCurrency(ctx *gin.Context) {

	oid, err := primitive.ObjectIDFromHex(ctx.Query("oid"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	currency, err := ctrl.CurrencyService.GetCurrency(oid)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, currency)
}

func (ctrl CurrencyController) UpdateCurrency(ctx *gin.Context) {

	var currency currencyModels.Currency

	if err := ctx.ShouldBindJSON(&currency); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.New("couldn't parse the request"))
		return
	}

	oid, err := primitive.ObjectIDFromHex(ctx.Query("oid"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errors.New("couldn't get an oid"))
		return
	}

	currency.ID = oid

	if err := ctrl.CurrencyService.UpdateCurrency(&currency); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true})
}

func (ctrl CurrencyController) GetAll(ctx *gin.Context) {

	currencyList, err := ctrl.CurrencyService.GetAll()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, currencyList)
}

func (ctrl CurrencyController) DeleteAll(ctx *gin.Context) {
	deleteCount, err := ctrl.CurrencyService.DeleteAll()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": deleteCount})
}
