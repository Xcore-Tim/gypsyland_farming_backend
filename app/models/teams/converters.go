package models

import (
	auth "gypsylandFarming/app/models/authentication"
	"strconv"
)

func (model *EditTeamAccessRequest) Convert() {
	auth.ConvertUserData(&model.UserData, model.UserIdentity)
	model.TeamEdit.TeamID, _ = strconv.Atoi(model.TeamEdit.ID)
}
