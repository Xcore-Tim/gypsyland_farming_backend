package services

import (
	"errors"
	accounts "gypsylandFarming/app/models/accounts"
	filters "gypsylandFarming/app/requests"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func (srvc AccountRequestServiceImpl) GetTLFPeindingRequests(requestBody *accounts.GetRequestBody, farmersPeindingRequests *[]accounts.FarmersPendingResponse) error {

	filter := filters.TeamleadFarmerRequestFilter(requestBody)
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

func (srvc AccountRequestServiceImpl) GetTLFInworkRequests(requestBody *accounts.GetRequestBody, farmersInworkRequests *[]accounts.FarmersInworkResponse) error {

	filter := filters.TeamleadFarmerRequestFilter(requestBody)
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

func (srvc AccountRequestServiceImpl) GetTLFCompletedRequests(requestBody *accounts.GetRequestBody, farmersCompletedResponse *[]accounts.FarmersCompletedResponse) error {

	filter := filters.TeamleadFarmerRequestFilter(requestBody)
	projection := filters.FarmerRequestProjection(requestBody)
	cursor, err := srvc.accountRequestTaskCollection.Find(srvc.ctx, filter, options.Find().SetProjection(projection))

	if err != nil {
		return err
	}

	for cursor.Next(srvc.ctx) {
		var accountRequest accounts.FarmersCompletedResponse

		if err := cursor.Decode(&accountRequest); err != nil {
			return errors.New("decode error")
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

func (srvc AccountRequestServiceImpl) GetTLFCancelledRequests(requestBody *accounts.GetRequestBody, farmersCancelledResponse *[]accounts.FarmersCancelledResponse) error {

	filter := filters.TeamleadFarmerRequestFilter(requestBody)
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
