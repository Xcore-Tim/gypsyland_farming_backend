package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (ctrl CurrencyController) GetCurrencyRates(ctx *gin.Context) {

	requestDate := time.Now().Format("02-01-2006")
	cr, err := ctrl.CurrencyRatesService.GetCurrencyRates(requestDate)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, cr)
}
