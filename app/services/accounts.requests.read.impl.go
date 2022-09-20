package services

import (
	"errors"
	"gypsyland_farming/app/models"
	db_requests "gypsyland_farming/app/requests"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (acc AccountRequestServiceImpl) GetTLFARequests(requestBody *models.GetRequestBody, accountRequestTasks *[]models.AccountRequestTask) error {

	filter := db_requests.TLFAdminRequest(requestBody)
	cursor, err := acc.accountRequestTaskCollection.Find(acc.ctx, filter)

	if err != nil {
		return err
	}

	for cursor.Next(acc.ctx) {
		var accountRequest models.AccountRequestTask

		if err := cursor.Decode(&accountRequest); err != nil {
			return err
		}
		*accountRequestTasks = append(*accountRequestTasks, accountRequest)
	}

	if err := cursor.Err(); err != nil {
		return err
	}

	cursor.Close(acc.ctx)

	if len(*accountRequestTasks) == 0 {
		return errors.New("documents not found")
	}

	return nil
}

func (acc AccountRequestServiceImpl) GetTeamleadRequests(requestBody *models.GetRequestBody, accountRequestTasks *[]models.AccountRequestTask) error {

	filter := db_requests.TeamleadRequest(requestBody)
	cursor, err := acc.accountRequestTaskCollection.Find(acc.ctx, filter)

	if err != nil {
		return err
	}

	for cursor.Next(acc.ctx) {
		var accountRequest models.AccountRequestTask

		if err := cursor.Decode(&accountRequest); err != nil {
			return err
		}
		*accountRequestTasks = append(*accountRequestTasks, accountRequest)
	}

	if err := cursor.Err(); err != nil {
		return err
	}

	cursor.Close(acc.ctx)

	if len(*accountRequestTasks) == 0 {
		return errors.New("documents not found")
	}

	return nil
}

func (acc AccountRequestServiceImpl) GetFarmerRequests(requestBody *models.GetRequestBody, accountRequestTasks *[]models.AccountRequestTask, teamAcess *models.TeamAccess) error {

	filter := db_requests.FarmerRequest(requestBody, *teamAcess)
	cursor, err := acc.accountRequestTaskCollection.Find(acc.ctx, filter)

	if err != nil {
		return err
	}

	for cursor.Next(acc.ctx) {
		var accountRequest models.AccountRequestTask

		if err := cursor.Decode(&accountRequest); err != nil {
			return err
		}
		*accountRequestTasks = append(*accountRequestTasks, accountRequest)
	}

	if err := cursor.Err(); err != nil {
		return err
	}

	cursor.Close(acc.ctx)

	if len(*accountRequestTasks) == 0 {
		return errors.New("documents not found")
	}

	return nil
}

func (acc AccountRequestServiceImpl) GetAll() ([]*models.AccountRequestTask, error) {

	var accountRequests []*models.AccountRequestTask

	cursor, err := acc.accountRequestTaskCollection.Find(acc.ctx, bson.D{{}})

	if err != nil {
		return nil, err
	}

	for cursor.Next(acc.ctx) {
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

	cursor.Close(acc.ctx)

	if len(accountRequests) == 0 {
		return nil, errors.New("documents not found")
	}

	return accountRequests, err
}

func (acc AccountRequestServiceImpl) AggregateFarmersData(groupedRepsonse *[]models.GroupedFarmersResponse) error {

	matchStage := bson.D{bson.E{Key: "$match", Value: bson.D{bson.E{Key: "status", Value: models.Complete}}}}

	groupStage := bson.D{
		{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$farmer"},
			{Key: "price", Value: bson.D{
				{Key: "$sum", Value: "$price"},
			}},
			{Key: "valid", Value: bson.D{
				{Key: "$sum", Value: "$valid"},
			}},
			{Key: "quantity", Value: bson.D{
				{Key: "$sum", Value: "$account_request.quantity"},
			}},
		}}}

	cursor, err := acc.accountRequestTaskCollection.Aggregate(acc.ctx, mongo.Pipeline{matchStage, groupStage})

	if err != nil {
		return err
	}

	for cursor.Next(acc.ctx) {
		var response models.GroupedFarmersResponse

		if err := cursor.Decode(&response); err != nil {
			return err
		}

		*groupedRepsonse = append(*groupedRepsonse, response)
	}

	return nil
}

func (acc AccountRequestServiceImpl) AggregateFarmersDataBSON() []bson.M {

	matchStage := bson.D{bson.E{Key: "$match", Value: bson.D{bson.E{Key: "status", Value: models.Complete}}}}

	groupStage := bson.D{
		{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$farmer"},
			{Key: "price", Value: bson.D{
				{Key: "$sum", Value: "$price"},
			}},
			{Key: "valid", Value: bson.D{
				{Key: "$sum", Value: "$valid"},
			}},
			{Key: "quantity", Value: bson.D{
				{Key: "$sum", Value: "$account_request.quantity"},
			}},
		}}}

	cursor, _ := acc.accountRequestTaskCollection.Aggregate(acc.ctx, mongo.Pipeline{matchStage, groupStage})

	var results []bson.M

	if err := cursor.All(acc.ctx, &results); err != nil {
		panic(err.Error())
	}

	return results
}

func (acc AccountRequestServiceImpl) AggregateBuyersData(teamlead_id int) []bson.M {

	matchStage := bson.D{bson.E{Key: "$match", Value: bson.D{
		bson.E{Key: "status", Value: models.Complete},
		bson.E{Key: "team.teamlead.id", Value: teamlead_id},
	}}}

	groupStage := bson.D{
		{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$buyer"},
			{Key: "price", Value: bson.D{
				{Key: "$sum", Value: "$price"},
			}},
			{Key: "valid", Value: bson.D{
				{Key: "$sum", Value: "$valid"},
			}},
			{Key: "quantity", Value: bson.D{
				{Key: "$sum", Value: "$account_request.quantity"},
			}},
			{Key: "team", Value: bson.D{
				{Key: "$first", Value: "$team"},
			}},
		}}}

	cursor, _ := acc.accountRequestTaskCollection.Aggregate(acc.ctx, mongo.Pipeline{matchStage, groupStage})

	var results []bson.M

	if err := cursor.All(acc.ctx, &results); err != nil {
		panic(err.Error())
	}

	return results
}

func (acc AccountRequestServiceImpl) AggregateTeamsData(groupedRepsonse *[]models.GroupedTeamsResponse) error {

	matchStage := bson.D{bson.E{Key: "$match", Value: bson.D{bson.E{Key: "status", Value: models.Complete}}}}

	groupStage := bson.D{
		{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$team"},
			{Key: "price", Value: bson.D{
				{Key: "$sum", Value: "$price"},
			}},
			{Key: "valid", Value: bson.D{
				{Key: "$sum", Value: "$valid"},
			}},
			{Key: "quantity", Value: bson.D{
				{Key: "$sum", Value: "$account_request.quantity"},
			}},
		}}}

	cursor, err := acc.accountRequestTaskCollection.Aggregate(acc.ctx, mongo.Pipeline{matchStage, groupStage})

	if err != nil {
		return err
	}

	for cursor.Next(acc.ctx) {
		var response models.GroupedTeamsResponse

		if err := cursor.Decode(&response); err != nil {
			return err
		}

		*groupedRepsonse = append(*groupedRepsonse, response)
	}

	return nil
}
