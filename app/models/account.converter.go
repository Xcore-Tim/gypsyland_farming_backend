package models

import (
	"strconv"
)

func ConvertUserData(userData *UserData, userIdentity UserIdentity) {

	userData.UserID, _ = strconv.Atoi(userIdentity.UserID)
	userData.TeamID, _ = strconv.Atoi(userIdentity.TeamID)
	userData.RoleID, _ = strconv.Atoi(userIdentity.RoleID)
	userData.Username = userIdentity.Username
	userData.Token = userIdentity.Token
}
