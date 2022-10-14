package services

import (
	"context"
	"errors"
	models "gypsylandFarming/app/models/currency"

	// "gypsylandFarming/app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (srvc CurrencyServiceImpl) CreateCurrency(currency *models.Currency) (primitive.ObjectID, error) {

	result, err := srvc.currencyCollection.InsertOne(srvc.ctx, currency)

	if err != nil {
		return primitive.NilObjectID, err
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)

	if !ok {
		return primitive.NilObjectID, errors.New("couldn't get oid")
	}

	return oid, nil
}

func (srvc CurrencyServiceImpl) UpdateCurrency(currency *models.Currency) error {

	filter := bson.D{bson.E{Key: "_id", Value: currency.ID}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "name", Value: currency.Name},
		bson.E{Key: "iso", Value: currency.ISO},
		bson.E{Key: "symbol", Value: currency.Symbol},
	}}}

	result, _ := srvc.currencyCollection.UpdateOne(srvc.ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("no currency found for update")
	}

	return nil
}

func (srvc CurrencyServiceImpl) GetCurrency(oid primitive.ObjectID) (*models.Currency, error) {

	var currency models.Currency

	if err := srvc.currencyCollection.FindOne(srvc.ctx, bson.D{bson.E{Key: "_id", Value: oid}}).Decode(&currency); err != nil {
		return &currency, errors.New("no currency found")
	}

	return &currency, nil
}

func (srvc CurrencyServiceImpl) GetBaseCurrency() (*models.Currency, error) {

	var currency models.Currency

	if err := srvc.currencyCollection.FindOne(srvc.ctx, bson.D{bson.E{Key: "iso", Value: "USD"}}).Decode(&currency); err != nil {
		return &currency, errors.New("no currency found")
	}

	return &currency, nil
}

func (srvc CurrencyServiceImpl) GetAll() ([]*models.Currency, error) {

	var currencySlice []*models.Currency

	cursor, err := srvc.currencyCollection.Find(srvc.ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(srvc.ctx)

	for cursor.Next(srvc.ctx) {
		var currency models.Currency

		if err := cursor.Decode(&currency); err != nil {
			return nil, err
		}

		currencySlice = append(currencySlice, &currency)
	}

	defer cursor.Close(srvc.ctx)

	if len(currencySlice) == 0 {
		return nil, errors.New("no currency found")
	}

	return currencySlice, nil
}

func (srvc CurrencyServiceImpl) DeleteAll() (int, error) {

	result, err := srvc.currencyCollection.DeleteMany(srvc.ctx, bson.D{})

	if err != nil {
		return 0, err
	}

	deleteCount := result.DeletedCount
	return int(deleteCount), nil
}
