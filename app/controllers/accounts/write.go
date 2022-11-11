package controllers

import (
	accounts "gypsylandFarming/app/models/accounts"
	"strconv"

	"time"

	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ctrl AccountRequestController) CreateAccountRequest(ctx *gin.Context) {

	var requestBody accounts.CreateAccountRequestBody

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}

	requestBody.Convert()

	var accountRequestTask accounts.AccountRequestTask

	if requestBody.AccountRequestBody.LocationID != "" {
		location, err := ctrl.LocationService.GetLocation(requestBody.AccountRequestData.LocationID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		accountRequestTask.AccountRequest.Location = *location
	}

	accountType, err := ctrl.AccountTypesService.GetType(requestBody.AccountRequestData.TypeID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	team, err := ctrl.TeamService.GetTeamByNum(requestBody.UserData.TeamID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accountRequestTask.AccountRequest.Type = *accountType
	accountRequestTask.AccountRequest.Quantity = requestBody.AccountRequestData.Quantity

	accountRequestTask.DateCreated = time.Now().Unix()
	accountRequestTask.Description = requestBody.AccountRequestBody.Description
	accountRequestTask.Price = ctrl.WriteAccountRequestService.RoundFloat(requestBody.AccountRequestData.Price, 2)
	total := float64(accountRequestTask.AccountRequest.Quantity) * accountRequestTask.Price
	accountRequestTask.TotalSum = ctrl.WriteAccountRequestService.RoundFloat(total, 2)

	accountRequestTask.Team = *team
	accountRequestTask.Buyer.ID = requestBody.UserData.UserID
	accountRequestTask.Buyer.Name = requestBody.UserData.Username
	accountRequestTask.Buyer.Position = requestBody.UserData.RoleID

	if currencyID, err := primitive.ObjectIDFromHex(requestBody.AccountRequestBody.Currency); err == nil {
		requestDate := time.Now().Format("02-01-2006")
		ctrl.SetRequestCurrency(currencyID, requestDate, &accountRequestTask)
	}

	if err := ctrl.WriteAccountRequestService.CreateAccountRequest(&accountRequestTask); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": accountRequestTask.ID})
}

func (ctrl AccountRequestController) UpdateRequest(ctx *gin.Context) {

	orderID, err := primitive.ObjectIDFromHex(ctx.Query("orderID"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	var accountRequestUpdate accounts.UpdateRequestBody

	if err := ctx.ShouldBindJSON(&accountRequestUpdate); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	accountRequestUpdate.RequestID = orderID
	accountRequestUpdate.Convert()

	originalAccountRequest, err := ctrl.ReadAccountRequestService.GetRequestTask(&accountRequestUpdate.RequestID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if accountRequestUpdate.UpdateBody.Location != "" {
		if location, err := ctrl.LocationService.GetLocationByName(accountRequestUpdate.UpdateBody.Location); err == nil {
			originalAccountRequest.AccountRequest.Location = *location
		}
	}

	if accountRequestUpdate.UpdateBody.AccountType != "" {
		if accountType, err := ctrl.AccountTypesService.GetTypeByName(accountRequestUpdate.UpdateBody.AccountType); err == nil {
			originalAccountRequest.AccountRequest.Type = *accountType
		}
	}

	if accountRequestUpdate.UpdateBody.Quantity != originalAccountRequest.AccountRequest.Quantity {
		originalAccountRequest.AccountRequest.Quantity = accountRequestUpdate.UpdateBody.Quantity
	}

	if accountRequestUpdate.UpdateBody.Price != originalAccountRequest.Price {
		originalAccountRequest.Price = ctrl.WriteAccountRequestService.RoundFloat(accountRequestUpdate.UpdateBody.Price, 2)
	}

	total := float64(originalAccountRequest.AccountRequest.Quantity) * originalAccountRequest.Price
	originalAccountRequest.TotalSum = ctrl.WriteAccountRequestService.RoundFloat(total, 2)

	if accountRequestUpdate.UpdateBody.Currency != "" {

		if currencyID, err := primitive.ObjectIDFromHex(accountRequestUpdate.UpdateBody.Currency); err == nil {
			updateCurrency, err := ctrl.CurrencyService.GetCurrency(currencyID)

			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if updateCurrency.ID != originalAccountRequest.Currency.ID {
				requestDate := time.Unix(originalAccountRequest.DateCreated, 0).Format("02-01-2006")
				ctrl.SetRequestCurrency(currencyID, requestDate, originalAccountRequest)
			}
		} else {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

	}

	if err := ctrl.WriteAccountRequestService.UpdateRequest(originalAccountRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (ctrl AccountRequestController) TakeAccountRequest(ctx *gin.Context) {

	var requestData accounts.TakeAccountRequest

	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	requestData.Convert()

	if err := ctrl.WriteAccountRequestService.TakeAccountRequest(&requestData); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (ctrl AccountRequestController) CancelAccountRequest(ctx *gin.Context) {

	var cancelRequest accounts.CancelAccountRequest

	err := ctx.ShouldBindJSON(&cancelRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	cancelRequest.Convert()

	err = ctrl.WriteAccountRequestService.CancelAccountRequest(&cancelRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (ctrl AccountRequestController) CompleteAccountRequest(ctx *gin.Context) {

	var accountRequestCompleted accounts.CompleteAccountRequest

	if err := ctx.ShouldBindJSON(&accountRequestCompleted); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accountRequestCompleted.Convert()

	accountRequest, err := ctrl.ReadAccountRequestService.GetRequestTask(&accountRequestCompleted.RequestID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if accountRequest.Price != accountRequestCompleted.OrderInfo.Price && accountRequestCompleted.OrderInfo.Price != 0 {
		accountRequest.Price = ctrl.WriteAccountRequestService.RoundFloat(accountRequestCompleted.OrderInfo.Price, 2)
		total := float64(accountRequest.AccountRequest.Quantity) * accountRequest.Price
		accountRequest.TotalSum = ctrl.WriteAccountRequestService.RoundFloat(total, 2)
	}

	accountRequest.Valid = accountRequestCompleted.OrderInfo.Valid
	accountRequest.Description = accountRequestCompleted.OrderInfo.Description

	if accountRequestCompleted.OrderInfo.Link != "" {
		accountRequest.DownloadLink = accountRequestCompleted.OrderInfo.Link
	}

	if accountRequestCompleted.OrderInfo.CurrencyID != "" {

		if currencyID, err := primitive.ObjectIDFromHex(accountRequestCompleted.OrderInfo.CurrencyID); err == nil {
			updateCurrency, err := ctrl.CurrencyService.GetCurrency(currencyID)

			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if updateCurrency.ID != accountRequest.Currency.ID {
				requestDate := time.Unix(accountRequest.DateCreated, 0).Format("02-01-2006")
				ctrl.SetRequestCurrency(currencyID, requestDate, accountRequest)
			}
		} else {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

	}

	if err := ctrl.WriteAccountRequestService.CompleteAccountRequest(accountRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (ctrl AccountRequestController) ReturnAccountRequest(ctx *gin.Context) {

	var requestData accounts.TakeAccountRequest

	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	requestData.Convert()

	if err := ctrl.WriteAccountRequestService.ReturnAccountRequest(&requestData.RequestID); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (ctrl AccountRequestController) SetRequestCurrency(currencyID primitive.ObjectID, requestDate string, accountRequestTask *accounts.AccountRequestTask) error {

	currency, err := ctrl.CurrencyService.GetCurrency(currencyID)

	if err != nil {
		return err
	}

	currencyRates, err := ctrl.CurrencyRatesService.GetCurrencyRates(requestDate)

	if err != nil {
		return err
	}

	currency.Value = currencyRates[currency.ISO]

	if currency.Value == 0 {
		currency.Value = 1
	}

	accountRequestTask.Currency = *currency

	switch accountRequestTask.Currency.ISO {
	case "USD":
		accountRequestTask.BaseCurrency = accountRequestTask.Currency
		accountRequestTask.BaseCurrency.Value = 1
		accountRequestTask.BaseTotal = accountRequestTask.TotalSum
	default:
		baseCurrency, _ := ctrl.CurrencyService.GetBaseCurrency()
		baseCurrency.Value = currencyRates[baseCurrency.ISO]
		accountRequestTask.BaseCurrency = *baseCurrency
		baseRate := accountRequestTask.Currency.Value / accountRequestTask.BaseCurrency.Value
		baseValue := accountRequestTask.TotalSum * baseRate
		accountRequestTask.BaseTotal = ctrl.CurrencyRatesService.RoundFloat(baseValue, 2)
	}

	return nil

}

func (ctrl AccountRequestController) Test(ctx *gin.Context) {

	// period := models.Period{}
	// period.StartISO = ctx.Query("start")
	// period.EndISO = ctx.Query("end")

	// date_format := "2006-01-02"

	// if period.StartISO == "" {
	// 	period.StartISO = "1970-01-01"
	// }

	// period.StartDate, _ = time.Parse(date_format, period.StartISO)

	// if period.EndISO != "" {
	// 	period.EndDate, _ = time.Parse(date_format, period.EndISO)
	// 	return
	// }

	// period.EndDate = time.Now()

	i, err := strconv.ParseInt("1666822678", 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	ctx.JSON(http.StatusAccepted, tm)

}
