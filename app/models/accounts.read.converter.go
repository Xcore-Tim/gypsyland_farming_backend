package models

func (r *GetRequestBody) Convert() {
	ConvertUserData(&r.UserData, r.UserIdentity)
	ConvertPeriod(&r.Period)
}
