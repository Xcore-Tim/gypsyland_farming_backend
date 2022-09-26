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

func (srvc AccountRequestServiceImpl) UpdateRequest(requestUpdate *models.UpdateAccountRequest) error {

	var accountRequestTask models.AccountRequestTask

	query := bson.D{bson.E{Key: "_id", Value: requestUpdate.RequestID}}

	err := srvc.accountRequestTaskCollection.FindOne(srvc.ctx, query).Decode(&accountRequestTask)

	if err != nil {
		return err
	}

	filter := bson.D{bson.E{Key: "_id", Value: accountRequestTask.AccountRequest.ID}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "quantity", Value: requestUpdate.AccountRequest.Quantity},
		bson.E{Key: "description", Value: requestUpdate.Description},
	}}}

	result, _ := srvc.accountRequestCollection.UpdateOne(srvc.ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("no matched documents found for update")
	}

	filter = bson.D{bson.E{Key: "_id", Value: requestUpdate.RequestID}}
	update = bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "accountRequest.quantity", Value: requestUpdate.AccountRequest.Quantity},
		bson.E{Key: "description", Value: requestUpdate.Description},
		bson.E{Key: "dateUpdated", Value: time.Now().Unix()},
	}}}

	result, _ = srvc.accountRequestTaskCollection.UpdateOne(srvc.ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("no matched documents found for update")
	}

	return nil
}

func (srvc AccountRequestServiceImpl) UpdateRequestNew(requestUpdate *models.UpdateAccountRequest) error {

	filter := bson.D{bson.E{Key: "_id", Value: requestUpdate.RequestID}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "accountRequest.quantity", Value: requestUpdate.AccountRequest.Quantity},
		bson.E{Key: "description", Value: requestUpdate.Description},
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

func (srvc AccountRequestServiceImpl) ReturnAccountRequest(request_id *primitive.ObjectID) (*models.AccountRequestTask, error) {
	var accountRequestTask models.AccountRequestTask

	filter := bson.D{bson.E{Key: "_id", Value: request_id}}

	err := srvc.accountRequestTaskCollection.FindOne(srvc.ctx, filter).Decode(&accountRequestTask)

	if err != nil {
		return nil, errors.New("no matched documents found to return to pending")
	}

	newAccountRequestTask := models.AccountRequestTask{
		AccountRequest: accountRequestTask.AccountRequest,
		Description:    accountRequestTask.Description,
		Buyer:          accountRequestTask.Buyer,
		Team:           accountRequestTask.Team,
	}

	return &newAccountRequestTask, nil
}

func (srvc AccountRequestServiceImpl) DeleteAccountRequest(id *primitive.ObjectID) error {

	var accountRequestTask models.AccountRequestTask

	taskFilter := bson.D{bson.E{Key: "_id", Value: id}}

	err := srvc.accountRequestTaskCollection.FindOne(srvc.ctx, taskFilter).Decode(&accountRequestTask)

	if err != nil {
		return errors.New("no matched tasks found to delete")
	}

	requestFilter := bson.D{bson.E{Key: "_id", Value: accountRequestTask.AccountRequest.ID}}

	requestResult, _ := srvc.accountRequestCollection.DeleteOne(srvc.ctx, requestFilter)

	if requestResult.DeletedCount != 1 {
		return errors.New("no matched requests found to delete")
	}

	taskResult, _ := srvc.accountRequestTaskCollection.DeleteOne(srvc.ctx, taskFilter)

	if taskResult.DeletedCount != 1 {
		return errors.New("no matched tasks found to delete")
	}

	return nil
}
