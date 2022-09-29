package services

import (
	"context"
	"encoding/json"
	"errors"
	"gypsyland_farming/app/models"
	"io"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (srvc TeamAccessServiceImpl) GetFarmersAccesses(farmerAccesses *[]models.FarmerAccess, userData *models.UserData) error {

	basepath := "https://g-identity-test.azurewebsites.net"
	endpoint := "/v1/Identity/users/byRole/"
	urlPath := basepath + endpoint + "6"

	bearer := "BEARER " + userData.Token

	request, err := http.NewRequest(http.MethodGet, urlPath, nil)

	if err != nil {
		return err
	}

	request.Header.Add("Authorization", bearer)

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	var result []user
	if err := json.Unmarshal([]byte(body), &result); err != nil {
		return err
	}

	for _, farmer := range result {
		var access models.FarmerAccess

		access.Farmer.ID = farmer.ID
		access.Farmer.Name = farmer.Username
		access.Farmer.Position = 6

		access.Teams = make([]int, 1)

		if teams, err := srvc.GetAccesses(farmer.ID); err == nil {
			access.Teams = teams.Teams
		}

		*farmerAccesses = append(*farmerAccesses, access)
	}

	return nil

}

func (srvc TeamAccessServiceImpl) GetAccesses(farmer int) (*models.TeamAccess, error) {

	var teamAccess models.TeamAccess

	err := srvc.teamAccessCollection.FindOne(srvc.ctx, bson.D{bson.E{Key: "employee", Value: farmer}}).Decode(&teamAccess)

	if err != nil {
		return nil, err
	}

	return &teamAccess, err
}

func (srvc TeamAccessServiceImpl) GetAccessByNum(teamAccess *models.TeamAccess, farmerID int) error {

	var access models.TeamAccess

	if err := srvc.teamAccessCollection.FindOne(srvc.ctx, bson.D{bson.E{Key: "employee", Value: farmerID}}).Decode(&access); err != nil {
		return err
	}

	for team := range access.Teams {
		teamAccess.Teams = append(teamAccess.Teams, team)
	}

	return nil
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

		result, err := srvc.teamAccessCollection.InsertOne(srvc.ctx, &newAccess)

		oid, ok := result.InsertedID.(primitive.ObjectID)

		if !ok {
			return &newAccess, err
		}

		newAccess.ID = oid

		return &newAccess, err
	}

	return teamAccess, err
}

func (srvc TeamAccessServiceImpl) UpdateAccess(teamAccess *models.TeamAccess) {

	filter := bson.D{bson.E{Key: "_id", Value: teamAccess.ID}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "employee", Value: teamAccess.Employee}, bson.E{Key: "teams", Value: teamAccess.Teams}}}}

	_, _ = srvc.teamAccessCollection.UpdateOne(srvc.ctx, filter, update)
}
