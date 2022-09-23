package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

type AuthRequestData struct {
	Token    Token            `json:"token"`
	Username string           `json:"username"`
	UserID   int              `json:"userID"`
	TeamID   int              `json:"teamID"`
	RoleID   int              `json:"roleID"`
	Meta     ResponseDataMeta `json:"meta"`
}

type ResponseDataMeta struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserName string             `json:"username" bson:"username"`
	Password string             `json:"password,omitempty" bson:"password"`
	Employee Employee           `json:"employee" bson:"employee"`
}

type UserIdentity struct {
	UserID   string `json:"userID"`
	RoleID   string `json:"roleID"`
	TeamID   string `json:"teamID"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

type UserData struct {
	UserID   int    `json:"userID"`
	RoleID   int    `json:"roleID"`
	TeamID   int    `json:"teamID"`
	Username string `json:"username"`
	Token    string `json:"token"`
}
