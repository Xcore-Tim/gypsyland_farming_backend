package services

import (
	"context"
	"errors"
	"gypsyland_farming/app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountTypesServiceImpl struct {
	accountTypesCollection *mongo.Collection
	ctx                    context.Context
}

func NewAccountTypesService(accountTypesCollection *mongo.Collection, ctx context.Context) AccountTypesService {

	return &AccountTypesServiceImpl{
		accountTypesCollection: accountTypesCollection,
	}
}

func (srvc AccountTypesServiceImpl) CreateAccountType(accountType *models.AccountType) error {
	_, err := srvc.accountTypesCollection.InsertOne(srvc.ctx, accountType)
	return err
}

func (srvc AccountTypesServiceImpl) GetAll() ([]*models.AccountType, error) {

	var accountTypes []*models.AccountType

	cursor, err := srvc.accountTypesCollection.Find(srvc.ctx, bson.D{bson.E{}})

	if err != nil {
		return nil, err
	}

	for cursor.Next(srvc.ctx) {
		var accountType models.AccountType
		err = cursor.Decode(&accountType)

		if err != nil {
			return nil, err
		}

		accountTypes = append(accountTypes, &accountType)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(srvc.ctx)

	if len(accountTypes) == 0 {
		return nil, errors.New("no documents found")
	}

	return accountTypes, err

}

func (srvc AccountTypesServiceImpl) GetType(accountTypeID primitive.ObjectID) (*models.AccountType, error) {

	var accountType models.AccountType

	query := bson.D{bson.E{Key: "_id", Value: accountTypeID}}

	err := srvc.accountTypesCollection.FindOne(srvc.ctx, query).Decode(&accountType)

	return &accountType, err
}
