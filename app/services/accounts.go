package services

import (
	"gypsyland_farming/app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (srvc AccountRequestServiceImpl) GetAccountRequestData(requestID *primitive.ObjectID) (*models.AccountRequest, error) {

	var accountRequest models.AccountRequest

	if err := srvc.accountRequestCollection.FindOne(srvc.ctx, bson.D{bson.E{Key: "_id", Value: requestID}}).Decode(&accountRequest); err != nil {
		return &accountRequest, err
	}
	return &accountRequest, nil
}

type ReadAccountRequestService interface {
	GetAll() ([]*models.AccountRequestTask, error)

	GetRequest(*primitive.ObjectID) (*models.AccountRequestTask, error)
	GetAccountRequestData(*primitive.ObjectID) (*models.AccountRequest, error)

	GetRequests(*models.GetRequestBody, *[]models.AccountRequestTask, models.GetFunctions) error

	GetFarmerPeindingRequests(*models.GetRequestBody, *[]models.FarmersPendingResponse, models.TeamAccess) error
	GetFarmerInworkRequests(*models.GetRequestBody, *[]models.FarmersInworkResponse, models.TeamAccess) error
	GetFarmerCompletedRequests(*models.GetRequestBody, *[]models.FarmersCompletedResponse, models.TeamAccess) error
	GetFarmerCancelledRequests(*models.GetRequestBody, *[]models.FarmersCancelledResponse, models.TeamAccess) error

	GetBuyerPendingRequests(*models.GetRequestBody, *[]models.BuyersPendingResponse) error
	GetBuyerInworkRequests(*models.GetRequestBody, *[]models.BuyersInworkResponse) error
	GetBuyerCompletedRequests(*models.GetRequestBody, *[]models.BuyersCompletedResponse) error
	GetBuyerCancelledRequests(*models.GetRequestBody, *[]models.BuyersCancelledResponse) error

	AggregateFarmersData(*models.GetRequestBody, *[]models.GroupedFarmersResponse) error
	AggregateTeamsData(*models.GetRequestBody, *[]models.GroupedTeamsResponse) error
	AggregateBuyersData(*models.GetRequestBody, *[]models.GroupedBuyersResponse, int) error
}

type WriteAccountRequestService interface {
	CreateAccountRequest(*models.AccountRequestTask) error
	UpdateRequest(*models.UpdateRequestBody) error
	UpdateDownloadLink(string, string) error

	TakeAccountRequest(*models.TakeAccountRequest) error
	CancelAccountRequest(*models.CancelAccountRequest) error
	CompleteAccountRequest(*models.CompleteAccountRequest) error
	ReturnAccountRequest(*primitive.ObjectID) error

	DeleteAccountRequest(*primitive.ObjectID) error
}

type AccountTypesService interface {
	CreateAccountType(*models.AccountType) error
	GetAll() ([]*models.AccountType, error)
	GetType(primitive.ObjectID) (*models.AccountType, error)
	GetTypeByName(string) (*models.AccountType, error)
}
