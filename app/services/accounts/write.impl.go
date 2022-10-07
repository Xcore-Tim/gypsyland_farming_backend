package services

import (
	"errors"
	global "gypsylandFarming/app/models"
	accounts "gypsylandFarming/app/models/accounts"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (srvc AccountRequestServiceImpl) CreateAccountRequest(accountRequestTask *accounts.AccountRequestTask) error {

	accountRequestTask.DateCreated = time.Now().Unix()

	result, err := srvc.accountRequestTaskCollection.InsertOne(srvc.ctx, accountRequestTask)

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		accountRequestTask.ID = oid
	}

	return err
}

func (srvc AccountRequestServiceImpl) UpdateRequest(requestUpdate *accounts.UCResponseBody) error {

	filter := bson.D{bson.E{Key: "_id", Value: requestUpdate.ID}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "accountRequest", Value: requestUpdate.AccountRequest},
		bson.E{Key: "price", Value: requestUpdate.Price},
		bson.E{Key: "totalSum", Value: requestUpdate.Total},
		bson.E{Key: "description", Value: requestUpdate.Description},
		bson.E{Key: "dateUpdated", Value: time.Now().Unix()},
	}}}

	result := srvc.accountRequestTaskCollection.FindOneAndUpdate(srvc.ctx, filter, update)

	if result.Err() != nil {
		return errors.New("no matched documents found for update")
	}

	return nil
}

func (srvc AccountRequestServiceImpl) TakeAccountRequest(requestData *accounts.TakeAccountRequest) error {

	filter := bson.D{bson.E{Key: "_id", Value: requestData.RequestID}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "farmer", Value: requestData.Farmer},
		bson.E{Key: "status", Value: accounts.Inwork},
	}}}

	result, _ := srvc.accountRequestTaskCollection.UpdateOne(srvc.ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("no matched documents found for update")
	}

	return nil
}

func (srvc AccountRequestServiceImpl) CancelAccountRequest(cancelRequest *accounts.CancelAccountRequest) error {

	filter := bson.D{bson.E{Key: "_id", Value: cancelRequest.RequestID}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "status", Value: accounts.Canceled},
		bson.E{Key: "denialReason", Value: cancelRequest.DenialReason},
		bson.E{Key: "cancelledBy", Value: cancelRequest.CancelledBy},
		bson.E{Key: "dateCancelled", Value: time.Now().Unix()},
	}}}

	result, _ := srvc.accountRequestTaskCollection.UpdateOne(srvc.ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("no matched documents found for update")
	}

	return nil
}

func (srvc AccountRequestServiceImpl) CompleteAccountRequest(accountRequestCompleted *accounts.CompleteAccountRequest) error {

	filter := bson.D{bson.E{Key: "_id", Value: accountRequestCompleted.RequestID}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "status", Value: accounts.Complete},
		bson.E{Key: "valid", Value: accountRequestCompleted.OrderInfo.Valid},
		bson.E{Key: "price", Value: accountRequestCompleted.OrderInfo.Price},
		bson.E{Key: "description", Value: accountRequestCompleted.OrderInfo.Description},
		bson.E{Key: "totalSum", Value: accountRequestCompleted.TotalSum},
		bson.E{Key: "downloadLink", Value: accountRequestCompleted.OrderInfo.Link},
		bson.E{Key: "dateFinished", Value: time.Now().Unix()},
	}}}

	result := srvc.accountRequestTaskCollection.FindOneAndUpdate(srvc.ctx, filter, update)

	if result.Err() != nil {
		return errors.New("no matched documents found for update")
	}

	return nil
}

func (srvc AccountRequestServiceImpl) ReturnAccountRequest(requestID *primitive.ObjectID) error {

	farmer := &global.Employee{
		ID:       0,
		Name:     "",
		Position: 0,
	}

	filter := bson.D{bson.E{Key: "_id", Value: requestID}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "status", Value: accounts.Pending},
		bson.E{Key: "farmer", Value: farmer},
	}}}

	result := srvc.accountRequestTaskCollection.FindOneAndUpdate(srvc.ctx, filter, update)

	if result.Err() != nil {
		return errors.New("no matched documents found for update")
	}

	return nil
}
