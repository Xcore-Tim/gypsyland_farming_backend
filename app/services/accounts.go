package services

import (
	"gypsyland_farming/app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReadAccountRequestService interface {
	GetAll() ([]*models.AccountRequestTask, error)

	GetRequest(*primitive.ObjectID) (*models.AccountRequestTask, error)

	GetTLFARequests(*models.GetRequestBody, *[]models.AccountRequestTask) error
	GetTeamleadRequests(*models.GetRequestBody, *[]models.AccountRequestTask) error
	GetRequests(*models.GetRequestBody, *[]models.AccountRequestTask, models.GetFunctions) error

	GetFarmerPeindingRequests(*models.GetRequestBody, *[]models.FarmersPendingResponse, models.TeamAccess) error
	GetFarmerInworkRequests(*models.GetRequestBody, *[]models.FarmersInworkResponse, models.TeamAccess) error
	GetFarmerCompletedRequests(*models.GetRequestBody, *[]models.FarmersCompletedResponse, models.TeamAccess) error
	GetFarmerCancelledRequests(*models.GetRequestBody, *[]models.FarmersCancelledResponse, models.TeamAccess) error

	GetBuyerPendingRequests(*models.GetRequestBody, *[]models.BuyersPendingResponse) error
	GetBuyerInworkRequests(*models.GetRequestBody, *[]models.BuyersInworkResponse) error
	GetBuyerCompletedRequests(*models.GetRequestBody, *[]models.BuyersCompletedResponse) error
	GetBuyerCancelledRequests(*models.GetRequestBody, *[]models.BuyersCancelledResponse) error

	GetTeamleadPendingRequests(*models.GetRequestBody, *[]models.BuyersPendingResponse) error
	GetTeamleadInworkRequests(*models.GetRequestBody, *[]models.BuyersInworkResponse) error
	GetTeamleadCompletedRequests(*models.GetRequestBody, *[]models.BuyersCompletedResponse) error
	GetTeamleadCancelledRequests(*models.GetRequestBody, *[]models.BuyersCancelledResponse) error

	AggregateFarmersData(*[]models.GroupedFarmersResponse) error
	AggregateTeamsData(*[]models.GroupedTeamsResponse) error
	AggregateBuyersData(teamlead_id int) []bson.M
}

type WriteAccountRequestService interface {
	CreateAccountRequest(*models.AccountRequestTask) error
	UpdateRequest(*models.UpdateAccountRequest) error
	UpdateRequestNew(*models.UpdateAccountRequest) error

	TakeAccountRequest(*models.TakeAccountRequest) error
	CancelAccountRequest(*models.CancelAccountRequest) error
	CompleteAccountRequest(*models.CompleteAccountRequest) error
	ReturnAccountRequest(*primitive.ObjectID) (*models.AccountRequestTask, error)

	DeleteAccountRequest(*primitive.ObjectID) error
}

type AccountTypesService interface {
	CreateAccountType(*models.AccountType) error
	GetAll() ([]*models.AccountType, error)
	GetType(primitive.ObjectID) (*models.AccountType, error)
}
