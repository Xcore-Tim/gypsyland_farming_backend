package services

import (
	"context"
	"errors"

	"gypsyland_farming/app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthServiceImpl struct {
	UserCollection     *mongo.Collection
	EmployeeCollection *mongo.Collection
	ctx                context.Context
}

func NewAuthService(userCollection *mongo.Collection, employeeCollection *mongo.Collection, ctx context.Context) AuthService {

	return &AuthServiceImpl{
		UserCollection:     userCollection,
		EmployeeCollection: employeeCollection,
		ctx:                ctx,
	}
}

func (aut AuthServiceImpl) CreateUser(user *models.User) error {

	result, err := aut.EmployeeCollection.InsertOne(aut.ctx, user.Employee)

	if err != nil {
		return err
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)

	if !ok {
		return err
	}

	user.ID = oid

	_, err = aut.UserCollection.InsertOne(aut.ctx, user)

	return err
}

func (aut AuthServiceImpl) UpdateUser(user *models.User) error {

	filter := bson.D{bson.E{Key: "_id", Value: user.Employee.ID}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "name", Value: user.Employee.Name},
		bson.E{Key: "position", Value: user.Employee.Position},
	}}}

	result, _ := aut.EmployeeCollection.UpdateOne(aut.ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("no documents found for update")
	}

	filter = bson.D{bson.E{Key: "_id", Value: user.ID}}
	update = bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "username", Value: user.UserName},
		bson.E{Key: "password", Value: user.Password},
		bson.E{Key: "employee", Value: user.Employee},
	}}}

	result, _ = aut.UserCollection.UpdateOne(aut.ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("no documents found for update")
	}

	return nil
}

func (auth AuthServiceImpl) GetUser(authRequest *models.AuthRequest) (*models.User, bool) {

	var user models.User

	filter := bson.D{
		bson.E{Key: "username", Value: authRequest.Username},
		bson.E{Key: "password", Value: authRequest.Password},
	}

	err := auth.UserCollection.FindOne(auth.ctx, filter).Decode(&user)

	if err != nil {
		return nil, false
	}

	return &user, true
}

func (auth AuthServiceImpl) GetAll() ([]*models.User, error) {

	var users []*models.User

	cursor, err := auth.UserCollection.Find(auth.ctx, bson.D{{}})

	if err != nil {
		return nil, err
	}

	for cursor.Next(auth.ctx) {

		var user models.User
		err := cursor.Decode(&user)

		if err != nil {
			return nil, err
		}

		users = append(users, &user)

	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(auth.ctx)

	return users, err
}
