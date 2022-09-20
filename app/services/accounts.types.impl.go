package services

import (
	"context"
	"errors"
	"gypsyland_farming/app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountTypesServiceImpl struct {
	accountTypeCollection *mongo.Collection
	ctx                   context.Context
}

func NewAccountTypesService(accountTypeCollection *mongo.Collection, ctx context.Context) AccountTypesService {

	return &AccountTypesServiceImpl{
		accountTypeCollection: accountTypeCollection,
	}
}

func (ats AccountTypesServiceImpl) CreateAccountType(accountType *models.AccountType) error {
	_, err := ats.accountTypeCollection.InsertOne(ats.ctx, accountType)
	return err
}

func (ats AccountTypesServiceImpl) GetAllAccountTypes() ([]*models.AccountType, error) {

	var accountTypes []*models.AccountType

	cursor, err := ats.accountTypeCollection.Find(ats.ctx, bson.D{bson.E{}})

	if err != nil {
		return nil, err
	}

	for cursor.Next(ats.ctx) {
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

	cursor.Close(ats.ctx)

	if len(accountTypes) == 0 {
		return nil, errors.New("no documents found")
	}

	return accountTypes, err

}
