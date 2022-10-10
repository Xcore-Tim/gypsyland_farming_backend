package services

import (
	"context"

	// "gypsylandFarming/app/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type CurrencyServiceImpl struct {
	currencyCollection *mongo.Collection
	ctx                context.Context
}

func NewCurrencyService(currencyCollection *mongo.Collection, ctx context.Context) CurrencyService {

	return &CurrencyServiceImpl{
		currencyCollection: currencyCollection,
		ctx:                ctx,
	}
}

func (srvc CurrencyServiceImpl) CreateCurrency() error {

	return nil
}
