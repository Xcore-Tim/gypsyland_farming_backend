package services

import (
	"context"
	accounts "gypsylandFarming/app/models/accounts"
	requests "gypsylandFarming/app/requests"
	"math"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AccountRequestServiceImpl struct {
	accountRequestTaskCollection *mongo.Collection
	ctx                          context.Context
}

func NewReadAccountRequestService(accountRequestTaskCollection *mongo.Collection, ctx context.Context) ReadAccountRequestService {

	return &AccountRequestServiceImpl{
		accountRequestTaskCollection: accountRequestTaskCollection,
		ctx:                          ctx,
	}
}

func NewWriteAccountRequestService(accountRequestTaskCollection *mongo.Collection, ctx context.Context) WriteAccountRequestService {

	return &AccountRequestServiceImpl{
		accountRequestTaskCollection: accountRequestTaskCollection,
		ctx:                          ctx,
	}
}

func (srvc AccountRequestServiceImpl) GetAccountRequestData(requestID *primitive.ObjectID) (*accounts.AccountRequestTask, error) {

	var accountRequest accounts.AccountRequestTask

	filter := bson.D{bson.E{Key: "_id", Value: requestID}}
	projection := requests.UCProjection()

	cursor, err := srvc.accountRequestTaskCollection.Find(srvc.ctx, filter, options.Find().SetProjection(projection))

	if err != nil {
		return &accountRequest, err
	}

	for cursor.Next(srvc.ctx) {
		if err := cursor.Decode(&accountRequest); err != nil {
			return &accountRequest, err
		}
	}

	return &accountRequest, nil
}

func (srvc AccountRequestServiceImpl) GetRequestTask(requestID *primitive.ObjectID) (*accounts.AccountRequestTask, error) {

	var accountRequestTask accounts.AccountRequestTask

	filter := bson.D{bson.E{Key: "_id", Value: requestID}}

	if err := srvc.accountRequestTaskCollection.FindOne(srvc.ctx, filter).Decode(&accountRequestTask); err != nil {
		return &accountRequestTask, err
	}

	return &accountRequestTask, nil
}

func (srvc AccountRequestServiceImpl) DeleteAccountRequest(oid primitive.ObjectID) error {

	filter := bson.D{bson.E{Key: "_id", Value: oid}}

	if _, err := srvc.accountRequestTaskCollection.DeleteOne(srvc.ctx, filter); err != nil {
		return err
	}

	return nil
}

func (srvc AccountRequestServiceImpl) DeleteAll() (int, error) {

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
