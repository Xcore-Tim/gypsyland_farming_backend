package controllers

import (
	service "gypsylandFarming/app/services/currency"

	"github.com/gin-gonic/gin"
)

type CurrencyController struct {
	CurrencyService      service.CurrencyService
	CurrencyRatesService service.CurrencyRatesService
}

func NewCurrencyController(currencyService service.CurrencyService, currencyRatesService service.CurrencyRatesService) CurrencyController {
	return CurrencyController{
		CurrencyService:      currencyService,
		CurrencyRatesService: currencyRatesService,
	}
}

func (ctrl CurrencyController) RegisterUserRoutes(rg *gin.RouterGroup) {
	currencyGroup := rg.Group("/currency")

	currencyGroup.POST("/create", ctrl.CreateCurrency)
	currencyGroup.POST("/update", ctrl.UpdateCurrency)

	getGroup := currencyGroup.Group("/get")
	getGroup.POST("/currency", ctrl.GetCurrency)
	getGroup.POST("/all", ctrl.GetAll)

	deleteGroup := currencyGroup.Group("/delete")
	deleteGroup.POST("/all", ctrl.DeleteAll)

	ratesGroup := currencyGroup.Group("/rates")

	ratesGetGroup := ratesGroup.Group("/get")
	ratesGetGroup.POST("/all", ctrl.GetCurrencyRates)
}
