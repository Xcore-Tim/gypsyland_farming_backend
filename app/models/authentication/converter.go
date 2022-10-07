package models

import (
	global "gypsylandFarming/app/models"
	"strconv"
	"time"
)

func ConvertUserData(userData *UserData, userIdentity UserIdentity) {

	userData.UserID, _ = strconv.Atoi(userIdentity.UserID)
	userData.TeamID, _ = strconv.Atoi(userIdentity.TeamID)
	userData.RoleID, _ = strconv.Atoi(userIdentity.RoleID)
	userData.Username = userIdentity.Username
	userData.Token = userIdentity.Token
}

func ConvertPeriod(period *global.Period) {

	if period.StartISO == "" {
		period.EndDate = time.Now()
	} else if period.EndISO == "" {
		period.EndDate = time.Now()
	} else {
		date_format := "2006-01-02"
		period.EndDate, _ = time.Parse(date_format, period.EndISO)
		period.StartDate, _ = time.Parse(date_format, period.StartISO)
	}
}
