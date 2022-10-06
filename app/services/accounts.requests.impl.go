package services

import (
	"context"
	"gypsyland_farming/app/models"
	"math"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountRequestServiceImpl struct {
	accountRequestCollection     *mongo.Collection
	accountRequestTaskCollection *mongo.Collection
	ctx                          context.Context
}

func NewReadAccountRequestService(accountRequestCollection *mongo.Collection, accountRequestTaskCollection *mongo.Collection, ctx context.Context) ReadAccountRequestService {

	return &AccountRequestServiceImpl{
		accountRequestCollection:     accountRequestCollection,
		accountRequestTaskCollection: accountRequestTaskCollection,
		ctx:                          ctx,
	}
}

func NewWriteAccountRequestService(accountRequestCollection *mongo.Collection, accountRequestTaskCollection *mongo.Collection, ctx context.Context) WriteAccountRequestService {

	return &AccountRequestServiceImpl{
		accountRequestCollection:     accountRequestCollection,
		accountRequestTaskCollection: accountRequestTaskCollection,
		ctx:                          ctx,
	}
}

func (srvc AccountRequestServiceImpl) GetAccountRequestData(requestID *primitive.ObjectID) (*models.AccountRequest, error) {

	var accountRequest models.AccountRequest

	if err := srvc.accountRequestCollection.FindOne(srvc.ctx, bson.D{bson.E{Key: "_id", Value: requestID}}).Decode(&accountRequest); err != nil {
		return &accountRequest, err
	}
	return &accountRequest, nil
}

func (srvc AccountRequestServiceImpl) GetRequestTask(requestID *primitive.ObjectID) (*models.AccountRequestTask, error) {

	var accountRequestTask models.AccountRequestTask

	filter := bson.D{bson.E{Key: "_id", Value: requestID}}

	if err := srvc.accountRequestTaskCollection.FindOne(srvc.ctx, filter).Decode(&accountRequestTask); err != nil {
		return &accountRequestTask, err
	}

	return &accountRequestTask, nil
}

func (srvc AccountRequestServiceImpl) DeleteAccountRequests() (int, error) {

	filter := bson.D{bson.E{}}

	requestResult, err := srvc.accountRequestCollection.DeleteMany(srvc.ctx, filter)

	if err != nil {
		return 0, err
	}

	return int(requestResult.DeletedCount), nil
}

func (srvc AccountRequestServiceImpl) DeleteAccountRequestTasks() (int, error) {

	filter := bson.D{bson.E{}}

	requestResult, err := srvc.accountRequestTaskCollection.DeleteMany(srvc.ctx, filter)

	if err != nil {
		return 0, err
	}

	return int(requestResult.DeletedCount), nil
}

func (srvc AccountRequestServiceImpl) RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
