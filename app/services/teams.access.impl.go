package services

import (
	"context"
	"errors"
	"gypsyland_farming/app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TeamAccessServiceImpl struct {
	teamAccessCollection *mongo.Collection
	ctx                  context.Context
}

func NewTeamAccessService(teamAccessCollection *mongo.Collection, ctx context.Context) TeamAccessService {

	return &TeamAccessServiceImpl{
		teamAccessCollection: teamAccessCollection,
		ctx:                  ctx,
	}

}

func (srvc TeamAccessServiceImpl) GetAllAccesses() ([]*models.TeamAccess, error) {

	var teamAccesses []*models.TeamAccess

	cursor, err := srvc.teamAccessCollection.Find(srvc.ctx, bson.D{bson.E{}})

	if err != nil {
		return nil, err
	}

	for cursor.Next(srvc.ctx) {
		var teamAccess models.TeamAccess

		err := cursor.Decode(&teamAccess)

		if err != nil {
			return nil, err
		}

		teamAccesses = append(teamAccesses, &teamAccess)

	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(srvc.ctx)

	if len(teamAccesses) == 0 {
		return nil, errors.New("documents not found")
	}

	return teamAccesses, err
}

func (srvc TeamAccessServiceImpl) GetAccesses(farmer int) (*models.TeamAccess, error) {

	var teamAccess models.TeamAccess

	err := srvc.teamAccessCollection.FindOne(srvc.ctx, bson.D{bson.E{Key: "employee", Value: farmer}}).Decode(&teamAccess)

	if err != nil {
		return nil, err
	}

	return &teamAccess, err
}

func (srvc TeamAccessServiceImpl) AddAccess(employee int, team int) error {

	teamAccess, err := srvc.CheckAccess(employee)

	if err != nil {
		return err
	}

	var found bool

	for _, v := range teamAccess.Teams {
		if v == team {
			found = true
		}
	}

	if !found {
		teamAccess.Teams = append(teamAccess.Teams, team)
		srvc.UpdateAccess(teamAccess)
	}

	return nil
}

func (srvc TeamAccessServiceImpl) RevokeAccess(employee int, team int) error {

	teamAccess, err := srvc.CheckAccess(employee)

	if err != nil {
		return err
	}

	for i, v := range teamAccess.Teams {
		if v == team {
			teamAccess.Teams = append(teamAccess.Teams[:i], teamAccess.Teams[i+1:]...)
			srvc.UpdateAccess(teamAccess)
		}
	}

	return nil
}

func (srvc TeamAccessServiceImpl) CheckAccess(employee int) (*models.TeamAccess, error) {

	var teamAccess *models.TeamAccess
	query := bson.D{bson.E{Key: "employee", Value: employee}}

	err := srvc.teamAccessCollection.FindOne(srvc.ctx, query).Decode(&teamAccess)

	if err != nil {
		var newAccess models.TeamAccess
		newAccess.Employee = employee

		_, err := srvc.teamAccessCollection.InsertOne(srvc.ctx, &newAccess)
		return &newAccess, err
	}

	return teamAccess, err
}

func (srvc TeamAccessServiceImpl) UpdateAccess(teamAccess *models.TeamAccess) {

	filter := bson.D{bson.E{Key: "_id", Value: teamAccess.ID}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "employee", Value: teamAccess.Employee}, bson.E{Key: "teams", Value: teamAccess.Teams}}}}

	_, _ = srvc.teamAccessCollection.UpdateOne(srvc.ctx, filter, update)
}
