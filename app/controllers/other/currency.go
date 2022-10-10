package controllers

type CurrencyController struct {
	CurrencyService string
}

func NewCurrencyController(currencyService string) CurrencyController {
	return CurrencyController{
		CurrencyService: currencyService,
	}
}
