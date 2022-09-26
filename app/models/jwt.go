package models

import "github.com/golang-jwt/jwt/v4"

type AuthCustomClaims struct {
	Name   string `json:"name"`
	IsUser bool   `json:"user"`
	// jwt.StandardClaims
	jwt.RegisteredClaims
}

type Token struct {
	AccessToken  string `json:"accessToken"`
	RequestToken string `json:"requestToken"`
}
