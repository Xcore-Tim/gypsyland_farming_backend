package services

import (
	"context"
	"errors"
	"gypsylandFarming/app/models"

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

func (srvc AccountTypesServiceImpl) GetTypeByName(name string) (*models.AccountType, error) {

	var accountType models.AccountType

	query := bson.D{bson.E{Key: "name", Value: name}}

	err := srvc.accountTypesCollection.FindOne(srvc.ctx, query).Decode(&accountType)

	return &accountType, err
}

func (srvc AccountTypesServiceImpl) DeleteType(oid *primitive.ObjectID) error {

	filter := bson.D{bson.E{Key: "_id", Value: oid}}
	result, _ := srvc.accountTypesCollection.DeleteOne(srvc.ctx, filter)

	if result.DeletedCount != 1 {
		return errors.New("no matched documents found for delete")
	}

	return nil
}

func (srvc AccountTypesServiceImpl) DeleteAll() (int, error) {

	filter := bson.D{bson.E{}}

	result, err := srvc.accountTypesCollection.DeleteMany(srvc.ctx, filter)

	if err != nil {
		return 0, err
	}

	return int(result.DeletedCount), nil

}
