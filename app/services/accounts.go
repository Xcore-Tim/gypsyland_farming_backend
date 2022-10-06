package services

import (
	"gypsyland_farming/app/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReadAccountRequestService interface {
	GetAll() ([]*models.AccountRequestTask, error)

	GetRequestTask(*primitive.ObjectID) (*models.AccountRequestTask, error)
	GetAccountRequestData(*primitive.ObjectID) (*models.AccountRequest, error)

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

	TakeAccountRequest(*models.TakeAccountRequest) error
	CancelAccountRequest(*models.CancelAccountRequest) error
	CompleteAccountRequest(*models.CompleteAccountRequest) error
	ReturnAccountRequest(*primitive.ObjectID) error

	DeleteAccountRequests() (int, error)
	DeleteAccountRequestTasks() (int, error)
	RoundFloat(float64, uint) float64
}

type AccountTypesService interface {
	CreateAccountType(*models.AccountType) error
	GetAll() ([]*models.AccountType, error)
	GetType(primitive.ObjectID) (*models.AccountType, error)
	GetTypeByName(string) (*models.AccountType, error)
}
