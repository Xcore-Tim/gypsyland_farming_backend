package services

import (
	"gypsyland_farming/app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReadAccountRequestService interface {
	GetAll() ([]*models.AccountRequestTask, error)
	GetTLFARequests(*models.GetRequestBody, *[]models.AccountRequestTask) error
	GetTeamleadRequests(*models.GetRequestBody, *[]models.AccountRequestTask) error
	GetFarmerRequests(*models.GetRequestBody, *[]models.AccountRequestTask, *models.TeamAccess) error

	AggregateFarmersData(*[]models.GroupedFarmersResponse) error
	AggregateTeamsData(*[]models.GroupedTeamsResponse) error
	AggregateBuyersData(teamlead_id int) []bson.M
	AggregateFarmersDataBSON() []bson.M
}

type WriteAccountRequestService interface {
	CreateAccountRequest(*models.AccountRequestTask) error
	UpdateAccountRequest(*models.AccountRequestTask) error
	UpdateRequest(*models.AccountRequestUpdate) error

	TakeAccountRequest(*models.Employee, *primitive.ObjectID) error
	CancelAccountRequest(*primitive.ObjectID, string) error
	CompleteAccountRequest(*models.AccountRequestCompleted) error
	ReturnAccountRequest(*primitive.ObjectID) (*models.AccountRequestTask, error)

	DeleteAccountRequest(*primitive.ObjectID) error
}

type AccountTypesService interface {
	CreateAccountType(*models.AccountType) error
	GetAllAccountTypes() ([]*models.AccountType, error)
}
