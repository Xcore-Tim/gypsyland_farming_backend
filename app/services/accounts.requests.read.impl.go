package services

import (
	"errors"
	"gypsyland_farming/app/models"
	filters "gypsyland_farming/app/requests"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (srvc AccountRequestServiceImpl) GetAll() ([]*models.AccountRequestTask, error) {

	var accountRequests []*models.AccountRequestTask

	cursor, err := srvc.accountRequestTaskCollection.Find(srvc.ctx, bson.D{{}})

	if err != nil {
		return nil, err
	}

	for cursor.Next(srvc.ctx) {
		var accountRequest models.AccountRequestTask
		err := cursor.Decode(&accountRequest)

		if err != nil {
			return nil, err
		}
		accountRequests = append(accountRequests, &accountRequest)

	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(srvc.ctx)

	if len(accountRequests) == 0 {
		return nil, errors.New("documents not found")
	}

	return accountRequests, err
}

func (srvc AccountRequestServiceImpl) GetRequest(requestID *primitive.ObjectID) (*models.AccountRequestTask, error) {

	var accountRequestTask models.AccountRequestTask

	filter := bson.D{bson.E{Key: "_id", Value: requestID}}

	if err := srvc.accountRequestTaskCollection.FindOne(srvc.ctx, filter).Decode(&accountRequestTask); err != nil {
		return &accountRequestTask, err
	}

	return &accountRequestTask, nil
}

func (srvc AccountRequestServiceImpl) GetTLFARequests(requestBody *models.GetRequestBody, accountRequestTasks *[]models.AccountRequestTask) error {

	filter := filters.TLFAdminRequest(requestBody)
	cursor, err := srvc.accountRequestTaskCollection.Find(srvc.ctx, filter)

	if err != nil {
		return err
	}

	for cursor.Next(srvc.ctx) {
		var accountRequest models.AccountRequestTask

		if err := cursor.Decode(&accountRequest); err != nil {
			return err
		}
		*accountRequestTasks = append(*accountRequestTasks, accountRequest)
	}

	if err := cursor.Err(); err != nil {
		return err
	}

	cursor.Close(srvc.ctx)

	if len(*accountRequestTasks) == 0 {
		return errors.New("documents not found")
	}

	return nil
}

func (srvc AccountRequestServiceImpl) GetRequests(requestBody *models.GetRequestBody, accountRequestTasks *[]models.AccountRequestTask, function models.GetFunctions) error {

	filter := function(requestBody)
	cursor, err := srvc.accountRequestTaskCollection.Find(srvc.ctx, filter)

	if err != nil {
		return err
	}

	for cursor.Next(srvc.ctx) {
		var accountRequest models.AccountRequestTask

		if err := cursor.Decode(&accountRequest); err != nil {
			return err
		}
		*accountRequestTasks = append(*accountRequestTasks, accountRequest)
	}

	if err := cursor.Err(); err != nil {
		return err
	}

	cursor.Close(srvc.ctx)

	if len(*accountRequestTasks) == 0 {
		return errors.New("documents not found")
	}

	return nil
}

func (srvc AccountRequestServiceImpl) AggregateFarmersData(groupedRepsonse *[]models.GroupedFarmersResponse) error {

	matchStage, groupStage := filters.AggregateFarmersData()

	cursor, err := srvc.accountRequestTaskCollection.Aggregate(srvc.ctx, mongo.Pipeline{matchStage, groupStage})

	if err != nil {
		return err
	}

	for cursor.Next(srvc.ctx) {
		var response models.GroupedFarmersResponse

		if err := cursor.Decode(&response); err != nil {
			return err
		}

		*groupedRepsonse = append(*groupedRepsonse, response)
	}

	return nil
}

func (srvc AccountRequestServiceImpl) AggregateBuyersData(teamlead_id int) []bson.M {

	matchStage, groupStage := filters.AggregateBuyersData(teamlead_id)

	cursor, _ := srvc.accountRequestTaskCollection.Aggregate(srvc.ctx, mongo.Pipeline{matchStage, groupStage})

	var results []bson.M

	if err := cursor.All(srvc.ctx, &results); err != nil {
		panic(err.Error())
	}

	return results
}

func (srvc AccountRequestServiceImpl) AggregateTeamsData(groupedRepsonse *[]models.GroupedTeamsResponse) error {

	matchStage, groupStage := filters.AggregateTeamsData()

	cursor, err := srvc.accountRequestTaskCollection.Aggregate(srvc.ctx, mongo.Pipeline{matchStage, groupStage})

	if err != nil {
		return err
	}

	for cursor.Next(srvc.ctx) {
		var response models.GroupedTeamsResponse

		if err := cursor.Decode(&response); err != nil {
			return err
		}

		*groupedRepsonse = append(*groupedRepsonse, response)
	}

	return nil
}
