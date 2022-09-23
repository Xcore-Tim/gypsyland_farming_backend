package services

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type AuthServiceImpl struct {
	UserCollection *mongo.Collection
	ctx            context.Context
}

func NewAuthService(userCollection *mongo.Collection, ctx context.Context) AuthService {

	return &AuthServiceImpl{
		UserCollection: userCollection,
		ctx:            ctx,
	}
}
