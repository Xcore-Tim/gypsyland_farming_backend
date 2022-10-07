package models

import (
	global "gypsylandFarming/app/models"
	auth "gypsylandFarming/app/models/authentication"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Team struct {
	ID       int             `json:"id" bson:"id"`
	Number   int             `json:"number" bson:"number"`
	TeamLead global.Employee `json:"teamlead" bson:"teamlead"`
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
	UserIdentity auth.UserIdentity `json:"userIdentity"`
	UserData     auth.UserData
	TeamEdit     TeamEdit
}

type FarmerAccess struct {
	Farmer global.Employee `json:"farmer"`
	Teams  []int           `json:"teams"`
}
