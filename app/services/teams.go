package services

import (
	"gypsyland_farming/app/models"
)

type TeamService interface {
	CreateTeam(*models.Team) error
	GetAllTeams() (*[]models.Team, error)
	GetDropdown(*models.TeamAccess, *models.EditTeamAccessRequest) (*[]models.TeamNumber, error)
	GetTLTeam(*models.Employee) (*models.Team, error)
	GetTeamByNum(int) (*models.Team, error)
	ImportTeams(string) error
}

type TeamAccessService interface {
	GetAllAccesses() ([]*models.TeamAccess, error)
	GetFarmersAccesses(*[]models.FarmerAccess, *models.UserData) error
	GetAccesses(int) (*models.TeamAccess, error)
	GetAccessByNum(*models.TeamAccess, int) error
	AddAccess(int, int) error
	RevokeAccess(int, int) error
}
