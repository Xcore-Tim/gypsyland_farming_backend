package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthResponse struct {
	Code int          `json:"code"`
	Data ResponseData `json:"data"`
}

type AuthRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type ResponseData struct {
	Token    Token            `json:"token"`
	FullName string           `json:"fullname"`
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
	UserID   int    `json:"userID"`
	RoleID   int    `json:"roleID"`
	TeamID   int    `json:"teamID"`
	FullName string `json:"fullname"`
	Token    string `json:"token"`
}
