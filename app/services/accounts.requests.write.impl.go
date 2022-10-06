package services

import (
	"errors"
	"gypsyland_farming/app/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (srvc AccountRequestServiceImpl) CreateAccountRequest(accountRequestTask *models.AccountRequestTask) error {

	result, err := srvc.accountRequestCollection.InsertOne(srvc.ctx, accountRequestTask.AccountRequest)

	if err != nil {
		return err
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)

	if !ok {
		return err
	}

	accountRequestTask.AccountRequest.ID = oid
	accountRequestTask.ID = oid
	accountRequestTask.DateCreated = time.Now().Unix()

	_, err = srvc.accountRequestTaskCollection.InsertOne(srvc.ctx, accountRequestTask)
	return err
}

func (srvc AccountRequestServiceImpl) UpdateRequest(requestUpdate *models.UpdateRequestBody) error {

	filter := bson.D{bson.E{Key: "_id", Value: requestUpdate.RequestID}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "accountRequest", Value: requestUpdate.AccountRequest},
		bson.E{Key: "description", Value: requestUpdate.UpdateBody.Description},
		bson.E{Key: "dateUpdated", Value: time.Now().Unix()},
	}}}

	result := srvc.accountRequestTaskCollection.FindOneAndUpdate(srvc.ctx, filter, update)

	if result.Err() != nil {
		return errors.New("no matched documents found for update")
	}

	return nil
}

func (srvc AccountRequestServiceImpl) TakeAccountRequest(requestData *models.TakeAccountRequest) error {

	filter := bson.D{bson.E{Key: "_id", Value: requestData.RequestID}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "farmer", Value: requestData.Farmer},
		bson.E{Key: "status", Value: models.Inwork},
	}}}

	result, _ := srvc.accountRequestTaskCollection.UpdateOne(srvc.ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("no matched documents found for update")
	}

	return nil
}

func (srvc AccountRequestServiceImpl) CancelAccountRequest(cancelRequest *models.CancelAccountRequest) error {

	filter := bson.D{bson.E{Key: "_id", Value: cancelRequest.RequestID}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "status", Value: models.Canceled},
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

func (srvc AccountRequestServiceImpl) CompleteAccountRequest(accountRequestCompleted *models.CompleteAccountRequest) error {

	filter := bson.D{bson.E{Key: "_id", Value: accountRequestCompleted.RequestID}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "status", Value: models.Complete},
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

	farmer := &models.Employee{
		ID:       0,
		Name:     "",
		Position: 0,
	}

	filter := bson.D{bson.E{Key: "_id", Value: requestID}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "status", Value: models.Pending},
		bson.E{Key: "farmer", Value: farmer},
	}}}

	result := srvc.accountRequestTaskCollection.FindOneAndUpdate(srvc.ctx, filter, update)

	if result.Err() != nil {
		return errors.New("no matched documents found for update")
	}

	return nil
}
