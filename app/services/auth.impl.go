package services

import (
	"context"
	"strconv"

	"bytes"
	"encoding/json"
	"gypsyland_farming/app/models"
	"io"
	"net/http"
)

type AuthServiceImpl struct {
	ctx context.Context
}

func NewAuthService(ctx context.Context) AuthService {
	return &AuthServiceImpl{
		ctx: ctx,
	}
}

func (srvc AuthServiceImpl) Login(authData *models.UserCredentials, authResponse *models.AuthResponseData) error {

	client := &http.Client{}

	urlPath := models.Basepath + models.EndpointAuth

	requestBody, _ := json.Marshal(map[string]string{
		"email":    authData.Email,
		"password": authData.Password,
	})

	bodyReader := bytes.NewBuffer(requestBody)

	postRequest, err := http.NewRequest(http.MethodPost, urlPath, bodyReader)

	if err != nil {
		return err
	}

	postRequest.Header.Set("Content-Type", models.ContentTypeAuth)

	resp, err := client.Do(postRequest)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	var userData models.UserData

	if err = json.Unmarshal([]byte(body), &userData); err != nil {
		return err
	}

	var user models.Username

	if err := srvc.GetFullname(&user, &userData); err != nil {
		return err
	}

	authResponse.Token.AccessToken = userData.Token
	authResponse.Token.RequestToken = userData.Token
	authResponse.Username = user.Username
	authResponse.UserID = userData.UserID
	authResponse.RoleID = userData.RoleID
	authResponse.TeamID = userData.TeamID

	return nil
}

func (srvc AuthServiceImpl) GetFullname(user *models.Username, userData *models.UserData) error {

	endpoint := "/v1/Identity/users/"
	urlPath := models.Basepath + endpoint + strconv.Itoa(userData.UserID)

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

	if err = json.Unmarshal([]byte(body), &user); err != nil {
		return err
	}

	return err
}

func (srvc AuthServiceImpl) AuthError(response *models.AuthResponseData, errorText string) {
	response.Meta.Error = errorText
	response.Meta.Message = "error"
}
