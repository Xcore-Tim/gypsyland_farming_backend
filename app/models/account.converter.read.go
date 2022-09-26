package models

func (r *GetRequestBody) Convert() {
	// r.UserData.UserID, _ = strconv.Atoi(r.UserIdentity.UserID)
	// r.UserData.TeamID, _ = strconv.Atoi(r.UserIdentity.TeamID)
	// r.UserData.RoleID, _ = strconv.Atoi(r.UserIdentity.RoleID)
	// r.UserData.Username = r.UserIdentity.Username
	// r.UserData.Token = r.UserIdentity.Token

	ConvertUserData(&r.UserData, r.UserIdentity)
}
