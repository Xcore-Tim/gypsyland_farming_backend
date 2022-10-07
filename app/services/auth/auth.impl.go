package services

import (
	"context"
	"strconv"

	"bytes"
	"encoding/json"
	auth "gypsylandFarming/app/models/authentication"
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

func (srvc AuthServiceImpl) Login(authData *auth.UserCredentials, authResponse *auth.AuthResponseData) error {

	urlPath := auth.Basepath + auth.EndpointAuth

	requestBody, _ := json.Marshal(map[string]string{
		"email":    authData.Email,
		"password": authData.Password,
	})

	bodyReader := bytes.NewBuffer(requestBody)

	postRequest, err := http.NewRequest(http.MethodPost, urlPath, bodyReader)

	if err != nil {
		return err
	}

	postRequest.Header.Set("Content-Type", auth.ContentTypeAuth)

	body, err := srvc.ClientRequest(postRequest)

	if err != nil {
		return err
	}

	var userData auth.UserData

	if err = json.Unmarshal([]byte(body), &userData); err != nil {
		return err
	}

	var user auth.Username

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

func (srvc AuthServiceImpl) GetFullname(user *auth.Username, userData *auth.UserData) error {

	endpoint := "/v1/Identity/users/"
	urlPath := auth.Basepath + endpoint + strconv.Itoa(userData.UserID)

	bearer := "BEARER " + userData.Token

	request, err := http.NewRequest(http.MethodGet, urlPath, nil)

	if err != nil {
		return err
	}

	request.Header.Add("Authorization", bearer)

	body, err := srvc.ClientRequest(request)

	if err != nil {
		return err
	}

	if err = json.Unmarshal([]byte(body), &user); err != nil {
		return err
	}

	return err
}

func (srvc AuthServiceImpl) AuthError(response *auth.AuthResponseData, errorText string) {
	response.Meta.Error = errorText
	response.Meta.Message = "error"
}

func (srvc AuthServiceImpl) ClientRequest(request *http.Request) ([]byte, error) {
	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		panic("error making a request")
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic("error reading body")
	}

	return body, nil
}
