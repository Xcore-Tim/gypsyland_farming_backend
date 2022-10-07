package services

import (
	"errors"
	accounts "gypsylandFarming/app/models/accounts"
	teams "gypsylandFarming/app/models/teams"
	filters "gypsylandFarming/app/requests"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func (srvc AccountRequestServiceImpl) GetFarmerPeindingRequests(requestBody *accounts.GetRequestBody, farmersPeindingRequests *[]accounts.FarmersPendingResponse, teamAccess teams.TeamAccess) error {

	filter := filters.FarmerRequestFilter(requestBody, teamAccess)
	projection := filters.FarmerRequestProjection(requestBody)

	cursor, err := srvc.accountRequestTaskCollection.Find(srvc.ctx, filter, options.Find().SetProjection(projection))

	if err != nil {
		return err
	}

	for cursor.Next(srvc.ctx) {
		var accountRequest accounts.FarmersPendingResponse

		if err := cursor.Decode(&accountRequest); err != nil {
			return err
		}
		*farmersPeindingRequests = append(*farmersPeindingRequests, accountRequest)
	}

	if err := cursor.Err(); err != nil {
		return err
	}

	cursor.Close(srvc.ctx)

	if len(*farmersPeindingRequests) == 0 {
		return errors.New("documents not found")
	}

	return nil
}

func (srvc AccountRequestServiceImpl) GetFarmerInworkRequests(requestBody *accounts.GetRequestBody, farmersInworkRequests *[]accounts.FarmersInworkResponse, teamAccess teams.TeamAccess) error {

	filter := filters.FarmerRequestFilter(requestBody, teamAccess)
	projection := filters.FarmerRequestProjection(requestBody)
	cursor, err := srvc.accountRequestTaskCollection.Find(srvc.ctx, filter, options.Find().SetProjection(projection))

	if err != nil {
		return err
	}

	for cursor.Next(srvc.ctx) {
		var accountRequest accounts.FarmersInworkResponse

		if err := cursor.Decode(&accountRequest); err != nil {
			return err
		}
		*farmersInworkRequests = append(*farmersInworkRequests, accountRequest)
	}

	if err := cursor.Err(); err != nil {
		return err
	}

	cursor.Close(srvc.ctx)

	if len(*farmersInworkRequests) == 0 {
		return errors.New("documents not found")
	}

	return nil
}

func (srvc AccountRequestServiceImpl) GetFarmerCompletedRequests(requestBody *accounts.GetRequestBody, farmersCompletedResponse *[]accounts.FarmersCompletedResponse, teamAccess teams.TeamAccess) error {

	filter := filters.FarmerRequestFilter(requestBody, teamAccess)
	projection := filters.FarmerRequestProjection(requestBody)
	cursor, err := srvc.accountRequestTaskCollection.Find(srvc.ctx, filter, options.Find().SetProjection(projection))

	if err != nil {
		return err
	}

	for cursor.Next(srvc.ctx) {
		var accountRequest accounts.FarmersCompletedResponse

		if err := cursor.Decode(&accountRequest); err != nil {
			return err
		}
		*farmersCompletedResponse = append(*farmersCompletedResponse, accountRequest)
	}

	if err := cursor.Err(); err != nil {
		return err
	}

	cursor.Close(srvc.ctx)

	if len(*farmersCompletedResponse) == 0 {
		return errors.New("documents not found")
	}

	return nil
}

func (srvc AccountRequestServiceImpl) GetFarmerCancelledRequests(requestBody *accounts.GetRequestBody, farmersCancelledResponse *[]accounts.FarmersCancelledResponse, teamAccess teams.TeamAccess) error {

	filter := filters.FarmerRequestFilter(requestBody, teamAccess)
	projection := filters.FarmerRequestProjection(requestBody)
	cursor, err := srvc.accountRequestTaskCollection.Find(srvc.ctx, filter, options.Find().SetProjection(projection))

	if err != nil {
		return err
	}

	for cursor.Next(srvc.ctx) {
		var accountRequest accounts.FarmersCancelledResponse

		if err := cursor.Decode(&accountRequest); err != nil {
			return err
		}
		*farmersCancelledResponse = append(*farmersCancelledResponse, accountRequest)
	}

	if err := cursor.Err(); err != nil {
		return err
	}

	cursor.Close(srvc.ctx)

	if len(*farmersCancelledResponse) == 0 {
		return errors.New("documents not found")
	}

	return nil
}
