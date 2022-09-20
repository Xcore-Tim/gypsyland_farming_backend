package services

import (
	"context"

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
