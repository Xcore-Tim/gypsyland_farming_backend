package services

import (
	"errors"
	accounts "gypsylandFarming/app/models/accounts"
	filters "gypsylandFarming/app/requests"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (srvc AccountRequestServiceImpl) GetAll() ([]*accounts.AccountRequestTask, error) {

	var accountRequests []*accounts.AccountRequestTask

	cursor, err := srvc.accountRequestTaskCollection.Find(srvc.ctx, bson.D{{}})

	if err != nil {
		return nil, err
	}

	for cursor.Next(srvc.ctx) {
		var accountRequest accounts.AccountRequestTask
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

func (srvc AccountRequestServiceImpl) AggregateFarmersData(requestBody *accounts.GetRequestBody, groupedRepsonse *[]accounts.GroupedFarmersResponse) error {

	matchStage, groupStage := filters.AggregateFarmersData(requestBody)

	cursor, err := srvc.accountRequestTaskCollection.Aggregate(srvc.ctx, mongo.Pipeline{matchStage, groupStage})

	if err != nil {
		return err
	}

	for cursor.Next(srvc.ctx) {
		var response accounts.GroupedFarmersResponse

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

func (srvc AccountRequestServiceImpl) AggregateBuyersData(requestBody *accounts.GetRequestBody, groupedRepsonse *[]accounts.GroupedBuyersResponse, teamleadID int) error {

	matchStage, groupStage := filters.AggregateBuyersData(requestBody, teamleadID)
	cursor, err := srvc.accountRequestTaskCollection.Aggregate(srvc.ctx, mongo.Pipeline{matchStage, groupStage})

	if err != nil {
		return err
	}

	for cursor.Next(srvc.ctx) {
		var response accounts.GroupedBuyersResponse
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

func (srvc AccountRequestServiceImpl) AggregateTeamsData(requestBody *accounts.GetRequestBody, groupedRepsonse *[]accounts.GroupedTeamsResponse) error {

	matchStage, groupStage := filters.AggregateTeamsData(requestBody)

	cursor, err := srvc.accountRequestTaskCollection.Aggregate(srvc.ctx, mongo.Pipeline{matchStage, groupStage})

	if err != nil {
		return err
	}

	for cursor.Next(srvc.ctx) {
		var response accounts.GroupedTeamsResponse

		if err := cursor.Decode(&response); err != nil {
			return err
		}

		*groupedRepsonse = append(*groupedRepsonse, response)
	}

	return nil
}
