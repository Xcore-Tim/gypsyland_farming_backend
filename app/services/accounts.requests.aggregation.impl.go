package services

import (
	"errors"
	"gypsyland_farming/app/models"
	filters "gypsyland_farming/app/requests"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
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

func (srvc AccountRequestServiceImpl) AggregateFarmersData(requestBody *models.GetRequestBody, groupedRepsonse *[]models.GroupedFarmersResponse) error {

	matchStage, groupStage := filters.AggregateFarmersData(requestBody)

	cursor, err := srvc.accountRequestTaskCollection.Aggregate(srvc.ctx, mongo.Pipeline{matchStage, groupStage})

	if err != nil {
		return err
	}

	for cursor.Next(srvc.ctx) {
		var response models.GroupedFarmersResponse

		if err := cursor.Decode(&response); err != nil {
			return err
		}
		response.UserIdentity.UserID = strconv.Itoa(response.Farmer.ID)
		response.UserIdentity.RoleID = strconv.Itoa(response.Farmer.Position)
		response.UserIdentity.Username = response.Farmer.Name
		response.UserIdentity.Token = requestBody.UserData.Token

		*groupedRepsonse = append(*groupedRepsonse, response)
	}

	return nil
}

func (srvc AccountRequestServiceImpl) AggregateBuyersData(requestBody *models.GetRequestBody, groupedRepsonse *[]models.GroupedBuyersResponse, teamleadID int) error {

	matchStage, groupStage := filters.AggregateBuyersData(requestBody, teamleadID)
	cursor, err := srvc.accountRequestTaskCollection.Aggregate(srvc.ctx, mongo.Pipeline{matchStage, groupStage})

	if err != nil {
		return err
	}

	for cursor.Next(srvc.ctx) {
		var response models.GroupedBuyersResponse
		if err := cursor.Decode(&response); err != nil {
			return err
		}
		response.UserData.UserID = response.Buyer.ID
		response.UserData.RoleID = response.Buyer.Position
		response.UserData.TeamID = response.Team.ID
		response.UserData.Username = response.Buyer.Name
		response.UserData.Token = requestBody.UserData.Token

		*groupedRepsonse = append(*groupedRepsonse, response)
	}

	return nil
}

func (srvc AccountRequestServiceImpl) AggregateTeamsData(requestBody *models.GetRequestBody, groupedRepsonse *[]models.GroupedTeamsResponse) error {

	matchStage, groupStage := filters.AggregateTeamsData(requestBody)

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
