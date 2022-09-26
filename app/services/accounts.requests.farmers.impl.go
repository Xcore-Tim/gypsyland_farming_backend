package services

import (
	"errors"
	"gypsyland_farming/app/models"
	filters "gypsyland_farming/app/requests"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func (srvc AccountRequestServiceImpl) GetFarmerPeindingRequests(requestBody *models.GetRequestBody, farmersPeindingRequests *[]models.FarmersPendingResponse, teamAccess models.TeamAccess) error {

	filter := filters.FarmerRequestFilter(requestBody, teamAccess)
	projection := filters.FarmerRequestProjection(requestBody)

	cursor, err := srvc.accountRequestTaskCollection.Find(srvc.ctx, filter, options.Find().SetProjection(projection))

	if err != nil {
		return err
	}

	for cursor.Next(srvc.ctx) {
		var accountRequest models.FarmersPendingResponse

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

func (srvc AccountRequestServiceImpl) GetFarmerInworkRequests(requestBody *models.GetRequestBody, farmersInworkRequests *[]models.FarmersInworkResponse, teamAccess models.TeamAccess) error {

	filter := filters.FarmerRequestFilter(requestBody, teamAccess)
	projection := filters.FarmerRequestProjection(requestBody)
	cursor, err := srvc.accountRequestTaskCollection.Find(srvc.ctx, filter, options.Find().SetProjection(projection))

	if err != nil {
		return err
	}

	for cursor.Next(srvc.ctx) {
		var accountRequest models.FarmersInworkResponse

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

func (srvc AccountRequestServiceImpl) GetFarmerCompletedRequests(requestBody *models.GetRequestBody, farmersCompletedResponse *[]models.FarmersCompletedResponse, teamAccess models.TeamAccess) error {

	filter := filters.FarmerRequestFilter(requestBody, teamAccess)
	projection := filters.FarmerRequestProjection(requestBody)
	cursor, err := srvc.accountRequestTaskCollection.Find(srvc.ctx, filter, options.Find().SetProjection(projection))

	if err != nil {
		return err
	}

	for cursor.Next(srvc.ctx) {
		var accountRequest models.FarmersCompletedResponse

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

func (srvc AccountRequestServiceImpl) GetFarmerCancelledRequests(requestBody *models.GetRequestBody, farmersCancelledResponse *[]models.FarmersCancelledResponse, teamAccess models.TeamAccess) error {

	filter := filters.FarmerRequestFilter(requestBody, teamAccess)
	projection := filters.FarmerRequestProjection(requestBody)
	cursor, err := srvc.accountRequestTaskCollection.Find(srvc.ctx, filter, options.Find().SetProjection(projection))

	if err != nil {
		return err
	}

	for cursor.Next(srvc.ctx) {
		var accountRequest models.FarmersCancelledResponse

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
