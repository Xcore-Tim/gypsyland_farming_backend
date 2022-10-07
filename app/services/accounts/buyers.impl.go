package services

import (
	"errors"
	accounts "gypsylandFarming/app/models/accounts"
	filters "gypsylandFarming/app/requests"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func (srvc AccountRequestServiceImpl) GetBuyerPendingRequests(requestBody *accounts.GetRequestBody, buyersPendingReponse *[]accounts.BuyersPendingResponse) error {

	filter := filters.BuyerRequestFilter(requestBody)
	projection := filters.BuyerRequestProjection(requestBody)

	cursor, err := srvc.accountRequestTaskCollection.Find(srvc.ctx, filter, options.Find().SetProjection(projection))

	if err != nil {
		return err
	}

	for cursor.Next(srvc.ctx) {
		var accountRequest accounts.BuyersPendingResponse

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

func (srvc AccountRequestServiceImpl) GetBuyerInworkRequests(requestBody *accounts.GetRequestBody, buyersInworkReponse *[]accounts.BuyersInworkResponse) error {

	filter := filters.BuyerRequestFilter(requestBody)
	projection := filters.BuyerRequestProjection(requestBody)
	cursor, err := srvc.accountRequestTaskCollection.Find(srvc.ctx, filter, options.Find().SetProjection(projection))

	if err != nil {
		return err
	}

	for cursor.Next(srvc.ctx) {
		var accountRequest accounts.BuyersInworkResponse

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

func (srvc AccountRequestServiceImpl) GetBuyerCompletedRequests(requestBody *accounts.GetRequestBody, buyersCompletedResponse *[]accounts.BuyersCompletedResponse) error {

	filter := filters.BuyerRequestFilter(requestBody)
	projection := filters.BuyerRequestProjection(requestBody)
	cursor, err := srvc.accountRequestTaskCollection.Find(srvc.ctx, filter, options.Find().SetProjection(projection))

	if err != nil {
		return err
	}

	for cursor.Next(srvc.ctx) {
		var accountRequest accounts.BuyersCompletedResponse

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

func (srvc AccountRequestServiceImpl) GetBuyerCancelledRequests(requestBody *accounts.GetRequestBody, buyersCancelledResponse *[]accounts.BuyersCancelledResponse) error {

	filter := filters.BuyerRequestFilter(requestBody)
	projection := filters.BuyerRequestProjection(requestBody)
	cursor, err := srvc.accountRequestTaskCollection.Find(srvc.ctx, filter, options.Find().SetProjection(projection))

	if err != nil {
		return err
	}

	for cursor.Next(srvc.ctx) {
		var accountRequest accounts.BuyersCancelledResponse

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
