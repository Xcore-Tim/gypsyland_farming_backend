package services

import (
	"context"
	"encoding/json"
	"errors"
	"gypsyland_farming/app/models"
	filters "gypsyland_farming/app/requests"
	"io"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TeamServiceImpl struct {
	teamCollection *mongo.Collection
	ctx            context.Context
}

type user struct {
	Username string
	ID       int
	Teamid   int
}

type teamNumber struct {
	Number int `json:"number"`
}

func NewTeamsService(teamCollection *mongo.Collection, ctx context.Context) TeamService {

	return &TeamServiceImpl{
		teamCollection: teamCollection,
		ctx:            ctx,
	}
}

func (srvc TeamServiceImpl) CreateTeam(team *models.Team) error {

	_, err := srvc.teamCollection.InsertOne(srvc.ctx, team)
	return err
}

func (srvc TeamServiceImpl) GetAllTeams() (*[]models.Team, error) {

	cursor, err := srvc.teamCollection.Find(srvc.ctx, bson.D{bson.E{}})

	if err != nil {
		return nil, err
	}

	var teamList []models.Team

	for cursor.Next(srvc.ctx) {
		var team models.Team
		err := cursor.Decode(&team)

		if err != nil {
			return nil, err
		}

		teamList = append(teamList, team)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(srvc.ctx)

	if len(teamList) == 0 {
		return nil, errors.New("no documents found")
	}

	return &teamList, err
}

func (srvc TeamServiceImpl) GetDropdown(teamAccess *models.TeamAccess, editAccessRequest *models.EditTeamAccessRequest) (*[]teamNumber, error) {

	filter := filters.TeamDropdown(teamAccess)
	projection := filters.TeamsProjection()
	cursor, err := srvc.teamCollection.Find(srvc.ctx, filter, options.Find().SetProjection(projection))

	if err != nil {
		return nil, err
	}

	var teamList []teamNumber

	for cursor.Next(srvc.ctx) {
		var team teamNumber
		err := cursor.Decode(&team)

		if err != nil {
			return nil, err
		}

		teamList = append(teamList, team)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(srvc.ctx)

	if len(teamList) == 0 {
		return nil, errors.New("no documents found")
	}

	return &teamList, err
}

func (srvc TeamServiceImpl) GetTLTeam(employee *models.Employee) (*models.Team, error) {

	var team *models.Team
	query := bson.D{bson.E{Key: "teamlead", Value: employee}}

	err := srvc.teamCollection.FindOne(srvc.ctx, query).Decode(&team)

	return team, err
}

func (srvc TeamServiceImpl) GetTeamByNum(num int) (*models.Team, error) {

	var team *models.Team
	query := bson.D{bson.E{Key: "number", Value: num}}

	err := srvc.teamCollection.FindOne(srvc.ctx, query).Decode(&team)

	return team, err
}

func (srvc TeamServiceImpl) ImportTeams(token string) error {

	basepath := "https://g-identity-test.azurewebsites.net"
	endpoint := "/v1/Identity/users/byRole/"
	urlPath := basepath + endpoint + "2"

	bearer := "BEARER " + token

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

	for _, teamlead := range result {
		var team models.Team

		team.ID = teamlead.Teamid
		team.Number = teamlead.Teamid
		team.TeamLead.ID = teamlead.ID
		team.TeamLead.Name = teamlead.Username
		team.TeamLead.Position = models.TeamLead

		if _, err := srvc.GetTeamByNum(team.Number); err != nil {
			if err := srvc.CreateTeam(&team); err != nil {
				return err
			}
		}
	}

	return err
}
