package services

import (
	"errors"
	"gypsyland_farming/app/models"
	filters "gypsyland_farming/app/requests"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func (srvc AccountRequestServiceImpl) GetTeamleadRequests(requestBody *models.GetRequestBody, accountRequestTasks *[]models.AccountRequestTask) error {

	filter := filters.TeamleadRequestFilter(requestBody)
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

func (srvc AccountRequestServiceImpl) GetTeamleadPendingRequests(requestBody *models.GetRequestBody, buyersPendingReponse *[]models.BuyersPendingResponse) error {

	filter := filters.TeamleadRequestFilter(requestBody)
	projection := filters.BuyerRequestProjection(requestBody)

	cursor, err := srvc.accountRequestTaskCollection.Find(srvc.ctx, filter, options.Find().SetProjection(projection))

	if err != nil {
		return err
	}

	for cursor.Next(srvc.ctx) {
		var accountRequest models.BuyersPendingResponse

		if err := cursor.Decode(&accountRequest); err != nil {
			return err
		}
		*buyersPendingReponse = append(*buyersPendingReponse, accountRequest)
	}

	if err := cursor.Err(); err != nil {
		return err
	}

	cursor.Close(srvc.ctx)

	if len(*buyersPendingReponse) == 0 {
		return errors.New("documents not found")
	}

	return nil
}

func (srvc AccountRequestServiceImpl) GetTeamleadInworkRequests(requestBody *models.GetRequestBody, buyersInworkReponse *[]models.BuyersInworkResponse) error {

	filter := filters.TeamleadRequestFilter(requestBody)
	projection := filters.BuyerRequestProjection(requestBody)
	cursor, err := srvc.accountRequestTaskCollection.Find(srvc.ctx, filter, options.Find().SetProjection(projection))

	if err != nil {
		return err
	}

	for cursor.Next(srvc.ctx) {
		var accountRequest models.BuyersInworkResponse

		if err := cursor.Decode(&accountRequest); err != nil {
			return err
		}
		*buyersInworkReponse = append(*buyersInworkReponse, accountRequest)
	}

	if err := cursor.Err(); err != nil {
		return err
	}

	cursor.Close(srvc.ctx)

	if len(*buyersInworkReponse) == 0 {
		return errors.New("documents not found")
	}

	return nil
}

func (srvc AccountRequestServiceImpl) GetTeamleadCompletedRequests(requestBody *models.GetRequestBody, buyersCompletedResponse *[]models.BuyersCompletedResponse) error {

	filter := filters.TeamleadRequestFilter(requestBody)
	projection := filters.BuyerRequestProjection(requestBody)
	cursor, err := srvc.accountRequestTaskCollection.Find(srvc.ctx, filter, options.Find().SetProjection(projection))

	if err != nil {
		return err
	}

	for cursor.Next(srvc.ctx) {
		var accountRequest models.BuyersCompletedResponse

		if err := cursor.Decode(&accountRequest); err != nil {
			return err
		}
		*buyersCompletedResponse = append(*buyersCompletedResponse, accountRequest)
	}

	if err := cursor.Err(); err != nil {
		return err
	}

	cursor.Close(srvc.ctx)

	if len(*buyersCompletedResponse) == 0 {
		return errors.New("documents not found")
	}

	return nil
}

func (srvc AccountRequestServiceImpl) GetTeamleadCancelledRequests(requestBody *models.GetRequestBody, buyersCancelledResponse *[]models.BuyersCancelledResponse) error {

	filter := filters.TeamleadRequestFilter(requestBody)
	projection := filters.BuyerRequestProjection(requestBody)
	cursor, err := srvc.accountRequestTaskCollection.Find(srvc.ctx, filter, options.Find().SetProjection(projection))

	if err != nil {
		return err
	}

	for cursor.Next(srvc.ctx) {
		var accountRequest models.BuyersCancelledResponse

		if err := cursor.Decode(&accountRequest); err != nil {
			return err
		}
		*buyersCancelledResponse = append(*buyersCancelledResponse, accountRequest)
	}

	if err := cursor.Err(); err != nil {
		return err
	}

	cursor.Close(srvc.ctx)

	if len(*buyersCancelledResponse) == 0 {
		return errors.New("documents not found")
	}

	return nil
}
