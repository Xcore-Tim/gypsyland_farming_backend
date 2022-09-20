package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Team struct {
	ID       int      `json:"id,omitempty" bson:"id,omitempty"`
	Number   int      `json:"number" bson:"number"`
	TeamLead Employee `json:"teamlead" bson:"teamlead"`
}

type TeamAccess struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Employee int                `json:"employee" bson:"employee"`
	Teams    []int              `json:"teams" bson:"teams"`
}

type BuyerTeam struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Employee Employee           `json:"employee" bson:"employee"`
	Team     Team               `json:"team" bson:"team"`
}