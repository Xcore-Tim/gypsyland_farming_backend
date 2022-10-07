package models

import (
	global "gypsylandFarming/app/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	Basepath        string = "https://g-identity-test.azurewebsites.net"
	EndpointAuth    string = "/v1/accounts/auth"
	ContentTypeAuth string = "application/json-patch+json"
)

type AuthResponseData struct {
	Token    Token            `json:"token"`
	Username string           `json:"username"`
	UserID   int              `json:"userID"`
	TeamID   int              `json:"teamID"`
	RoleID   int              `json:"roleID"`
	Meta     ResponseMetaData `json:"meta"`
}

type ResponseMetaData struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserName string             `json:"username" bson:"username"`
	Password string             `json:"password,omitempty" bson:"password"`
	Employee global.Employee    `json:"employee" bson:"employee"`
}

type UserCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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

type Username struct {
	Username string
}
