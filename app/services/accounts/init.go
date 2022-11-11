package services

import (
	accounts "gypsylandFarming/app/models/accounts"
	teams "gypsylandFarming/app/models/teams"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReadAccountRequestService interface {
	GetAll() ([]*accounts.AccountRequestTask, error)

	GetRequestTask(*primitive.ObjectID) (*accounts.AccountRequestTask, error)
	GetAccountRequestData(*primitive.ObjectID) (*accounts.AccountRequestTask, error)

	GetFarmerPeindingRequests(*accounts.GetRequestBody, *[]accounts.FarmersPendingResponse, teams.TeamAccess) error
	GetFarmerInworkRequests(*accounts.GetRequestBody, *[]accounts.FarmersInworkResponse, teams.TeamAccess) error
	GetFarmerCompletedRequests(*accounts.GetRequestBody, *[]accounts.FarmersCompletedResponse, teams.TeamAccess) error
	GetFarmerCancelledRequests(*accounts.GetRequestBody, *[]accounts.FarmersCancelledResponse, teams.TeamAccess) error

	GetBuyerPendingRequests(*accounts.GetRequestBody, *[]accounts.BuyersPendingResponse) error
	GetBuyerInworkRequests(*accounts.GetRequestBody, *[]accounts.BuyersInworkResponse) error
	GetBuyerCompletedRequests(*accounts.GetRequestBody, *[]accounts.BuyersCompletedResponse) error
	GetBuyerCancelledRequests(*accounts.GetRequestBody, *[]accounts.BuyersCancelledResponse) error

	GetTLFPeindingRequests(*accounts.GetRequestBody, *[]accounts.FarmersPendingResponse) error
	GetTLFInworkRequests(*accounts.GetRequestBody, *[]accounts.FarmersInworkResponse) error
	GetTLFCompletedRequests(*accounts.GetRequestBody, *[]accounts.FarmersCompletedResponse) error
	GetTLFCancelledRequests(*accounts.GetRequestBody, *[]accounts.FarmersCancelledResponse) error

	AggregateFarmersData(*accounts.GetRequestBody, *[]accounts.GroupedFarmersResponse) error
	AggregateTeamsData(*accounts.GetRequestBody, *[]accounts.GroupedTeamsResponse) error
	AggregateBuyersData(*accounts.GetRequestBody, *[]accounts.GroupedBuyersResponse, int) error
}

type WriteAccountRequestService interface {
	CreateAccountRequest(*accounts.AccountRequestTask) error
	UpdateRequest(*accounts.AccountRequestTask) error

	TakeAccountRequest(*accounts.TakeAccountRequest) error
	CancelAccountRequest(*accounts.CancelAccountRequest) error
	CompleteAccountRequest(*accounts.AccountRequestTask) error
	ReturnAccountRequest(*primitive.ObjectID) error

	DeleteAccountRequest(primitive.ObjectID) error
	DeleteAll() (int, error)
	RoundFloat(float64, uint) float64
}
