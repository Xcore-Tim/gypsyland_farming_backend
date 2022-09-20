package services

import (
	"gypsyland_farming/app/models"
)

type TeamService interface {
	CreateTeam(*models.Team) error
	GetAllTeams() ([]*models.Team, error)
	GetTLTeam(*models.Employee) (*models.Team, error)
	GetTeamByNum(int) (*models.Team, error)
	ImportTeams(string) error
	ImportTeams1(string) string
}

type TeamAccessService interface {
	GetAllAccesses() ([]*models.TeamAccess, error)
	GetAccesses(int) (*models.TeamAccess, error)
	AddAccess(int, int) error
	RevokeAccess(int, int) error
}
