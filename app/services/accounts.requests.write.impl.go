package services

import (
	"errors"
	"gypsyland_farming/app/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (srv AccountRequestServiceImpl) CreateAccountRequest(accountRequestTask *models.AccountRequestTask) error {

	result, err := srv.accountRequestCollection.InsertOne(srv.ctx, accountRequestTask.AccountRequest)

	if err != nil {
		return err
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)

	if !ok {
		return err
	}

	accountRequestTask.AccountRequest.ID = oid

	accountRequestTask.DateCreated = time.Now().Unix()
	_, err = srv.accountRequestTaskCollection.InsertOne(srv.ctx, accountRequestTask)
	return err
}

func (acc AccountRequestServiceImpl) UpdateAccountRequest(accountRequest *models.AccountRequestTask) error {

	filter := bson.D{bson.E{Key: "_id", Value: accountRequest.AccountRequest.ID}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "type", Value: accountRequest.AccountRequest.Type},
		bson.E{Key: "location", Value: accountRequest.AccountRequest.Location},
		bson.E{Key: "quantity", Value: accountRequest.AccountRequest.Quantity},
	}}}

	result, _ := acc.accountRequestCollection.UpdateOne(acc.ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("no matched documents found for update")
	}

	filter = bson.D{bson.E{Key: "_id", Value: accountRequest.ID}}
	update = bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "account_request", Value: accountRequest.AccountRequest},
		bson.E{Key: "status", Value: accountRequest.Status},
		bson.E{Key: "buyer", Value: accountRequest.Buyer},
		bson.E{Key: "farmer", Value: accountRequest.Farmer},
		bson.E{Key: "team", Value: accountRequest.Team},
		bson.E{Key: "valid", Value: accountRequest.Valid},
		bson.E{Key: "price", Value: accountRequest.Price},
		bson.E{Key: "currency", Value: accountRequest.Currency},
		bson.E{Key: "date_updated", Value: time.Now().Unix()},
		bson.E{Key: "denial_reason", Value: accountRequest.DenialReason},
	}}}

	result, _ = acc.accountRequestTaskCollection.UpdateOne(acc.ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("no matched documents found for update")
	}

	return nil
}

func (acc AccountRequestServiceImpl) UpdateRequest(requestUpdate *models.AccountRequestUpdate) error {

	var accountRequestTask models.AccountRequestTask

	query := bson.D{bson.E{Key: "_id", Value: requestUpdate.ID}}

	err := acc.accountRequestTaskCollection.FindOne(acc.ctx, query).Decode(&accountRequestTask)

	if err != nil {
		return err
	}

	filter := bson.D{bson.E{Key: "_id", Value: accountRequestTask.AccountRequest.ID}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "quantity", Value: requestUpdate.Quantity},
		bson.E{Key: "price", Value: requestUpdate.Price},
		bson.E{Key: "description", Value: requestUpdate.Description},
	}}}

	result, _ := acc.accountRequestCollection.UpdateOne(acc.ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("no matched documents found for update")
	}

	filter = bson.D{bson.E{Key: "_id", Value: requestUpdate.ID}}
	update = bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "account_request.quantity", Value: requestUpdate.Quantity},
		bson.E{Key: "price", Value: requestUpdate.Price},
		bson.E{Key: "description", Value: requestUpdate.Description},
		bson.E{Key: "denial_reason", Value: requestUpdate.DenialReason},
	}}}

	result, _ = acc.accountRequestTaskCollection.UpdateOne(acc.ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("no matched documents found for update")
	}

	return nil
}

func (acc AccountRequestServiceImpl) TakeAccountRequest(farmer *models.Employee, request_id *primitive.ObjectID) error {

	filter := bson.D{bson.E{Key: "_id", Value: request_id}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "farmer", Value: farmer},
		bson.E{Key: "farmerID", Value: farmer.ID},
		bson.E{Key: "status", Value: models.Inwork},
	}}}

	result, _ := acc.accountRequestTaskCollection.UpdateOne(acc.ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("no matched documents found for update")
	}

	return nil
}

func (acc AccountRequestServiceImpl) CancelAccountRequest(request_id *primitive.ObjectID, description string) error {

	filter := bson.D{bson.E{Key: "_id", Value: request_id}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "status", Value: models.Canceled},
		bson.E{Key: "denial_reason", Value: description},
	}}}

	result, _ := acc.accountRequestTaskCollection.UpdateOne(acc.ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("no matched documents found for update")
	}

	return nil
}

func (acc AccountRequestServiceImpl) CompleteAccountRequest(accountRequestCompleted *models.AccountRequestCompleted) error {
	filter := bson.D{bson.E{Key: "_id", Value: accountRequestCompleted.ID}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "status", Value: models.Complete},
		bson.E{Key: "valid", Value: accountRequestCompleted.Valid},
		bson.E{Key: "price", Value: accountRequestCompleted.Price},
		bson.E{Key: "description", Value: accountRequestCompleted.Description},
		bson.E{Key: "total_sum", Value: accountRequestCompleted.TotalSum},
	}}}

	result, _ := acc.accountRequestTaskCollection.UpdateOne(acc.ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("no matched documents found for update")
	}

	return nil
}

func (acc AccountRequestServiceImpl) ReturnAccountRequest(request_id *primitive.ObjectID) (*models.AccountRequestTask, error) {
	var accountRequestTask models.AccountRequestTask

	filter := bson.D{bson.E{Key: "_id", Value: request_id}}

	err := acc.accountRequestTaskCollection.FindOne(acc.ctx, filter).Decode(&accountRequestTask)

	if err != nil {
		return nil, errors.New("no matched documents found to return to pending")
	}

	newAccountRequestTask := models.AccountRequestTask{
		AccountRequest: accountRequestTask.AccountRequest,
		Description:    accountRequestTask.Description,
		Buyer:          accountRequestTask.Buyer,
		BuyerID:        accountRequestTask.BuyerID,
		Team:           accountRequestTask.Team,
		TeamID:         accountRequestTask.TeamID,
	}

	return &newAccountRequestTask, nil
}

func (acc AccountRequestServiceImpl) DeleteAccountRequest(id *primitive.ObjectID) error {

	var accountRequestTask models.AccountRequestTask

	taskFilter := bson.D{bson.E{Key: "_id", Value: id}}

	err := acc.accountRequestTaskCollection.FindOne(acc.ctx, taskFilter).Decode(&accountRequestTask)

	if err != nil {
		return errors.New("no matched tasks found to delete")
	}

	requestFilter := bson.D{bson.E{Key: "_id", Value: accountRequestTask.AccountRequest.ID}}

	requestResult, _ := acc.accountRequestCollection.DeleteOne(acc.ctx, requestFilter)

	if requestResult.DeletedCount != 1 {
		return errors.New("no matched requests found to delete")
	}

	taskResult, _ := acc.accountRequestTaskCollection.DeleteOne(acc.ctx, taskFilter)

	if taskResult.DeletedCount != 1 {
		return errors.New("no matched tasks found to delete")
	}

	return nil
}
