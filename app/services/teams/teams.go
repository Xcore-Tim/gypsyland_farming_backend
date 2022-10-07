package teamService

import (
	global "gypsylandFarming/app/models"
	auth "gypsylandFarming/app/models/authentication"
	teams "gypsylandFarming/app/models/teams"
)

type TeamService interface {
	CreateTeam(*teams.Team) error
	GetAllTeams() (*[]teams.Team, error)
	GetDropdown(*teams.TeamAccess, *teams.EditTeamAccessRequest) (*[]teamNumber, error)
	GetTLTeam(*global.Employee) (*teams.Team, error)
	GetTeamByNum(int) (*teams.Team, error)
	ImportTeams(string) error
}

type TeamAccessService interface {
	GetAllAccesses() ([]*teams.TeamAccess, error)
	GetFarmersAccesses(*[]teams.FarmerAccess, *auth.UserData) error
	GetAccess(int) (*teams.TeamAccess, error)
	GetAccessByNum(*teams.TeamAccess, int) error
	AddAccess(int, int) error
	RevokeAccess(int, int) error
}
