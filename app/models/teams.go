package models

import (
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Team struct {
	ID       int      `json:"id" bson:"id"`
	Number   int      `json:"number" bson:"number"`
	TeamLead Employee `json:"teamlead" bson:"teamlead"`
}

type TeamNumber struct {
	Number int `json:"number"`
}

type TeamAccess struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Employee int                `json:"employee" bson:"employee"`
	Teams    []int              `json:"teams" bson:"teams"`
}

type TeamEdit struct {
	ID     string `json:"teamID"`
	TeamID int
}

type EditTeamAccessRequest struct {
	UserIdentity UserIdentity `json:"userIdentity"`
	UserData     UserData
	TeamEdit     TeamEdit
}

type EditTeamAccessRequestBackup struct {
	UserID   string `json:"userID"`
	Token    string `json:"token"`
	FarmerID int
}

type FarmerAccess struct {
	Farmer Employee `json:"farmer"`
	Teams  []int    `json:"teams"`
}

func (r *EditTeamAccessRequest) Convert() {
	ConvertUserData(&r.UserData, r.UserIdentity)
	r.TeamEdit.TeamID, _ = strconv.Atoi(r.TeamEdit.ID)
}
