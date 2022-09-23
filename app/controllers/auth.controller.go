package controllers

import (
	"gypsyland_farming/app/models"
	"gypsyland_farming/app/services"

	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userResponse struct {
	Username string
}

type AuthController struct {
	AuthService services.AuthService
	JWTService  services.JWTService
	TeamService services.TeamService
}

type AuthRequestData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewAuthController(authService services.AuthService, jwtService services.JWTService, teamService services.TeamService) AuthController {
	return AuthController{
		AuthService: authService,
		JWTService:  jwtService,
		TeamService: teamService,
	}
}

func (ctrl AuthController) Login(ctx *gin.Context) {

	var authResponse models.AuthRequestData
	client := &http.Client{}

	endpoint := "/v1/accounts/auth"
	urlPath := models.Basepath + endpoint

	var authData AuthRequestData

	if err := ctx.ShouldBindJSON(&authData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	requestBody, _ := json.Marshal(map[string]string{
		"email":    authData.Email,
		"password": authData.Password,
	})

	bodyReader := bytes.NewBuffer(requestBody)

	contentType := "application/json-patch+json"

	postRequest, err := http.NewRequest(http.MethodPost, urlPath, bodyReader)

	if err != nil {
		DenyAuthentication(&authResponse, err, ctx)
		return
	}

	postRequest.Header.Set("Content-Type", contentType)

	resp, err := client.Do(postRequest)

	if err != nil {
		DenyAuthentication(&authResponse, err, ctx)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		DenyAuthentication(&authResponse, err, ctx)
		return
	}

	var userData models.UserData

	if err = json.Unmarshal([]byte(body), &userData); err != nil {
		DenyAuthentication(&authResponse, err, ctx)
		return
	}

	var userResponse userResponse

	if err := GetFullname(&userResponse, &userData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authResponse.Token.AccessToken = userData.Token
	authResponse.Token.RequestToken = userData.Token
	authResponse.Username = userResponse.Username
	authResponse.UserID = userData.UserID
	authResponse.RoleID = userData.RoleID
	authResponse.TeamID = userData.TeamID

	ctx.JSON(http.StatusOK, authResponse)
}

func (ctrl AuthController) RegisterUserRoutes(rg *gin.RouterGroup) {

	authRequestGroup := rg.Group("/auth")

	authRequestGroup.POST("", ctrl.Login)

}

func DenyAuthentication(authReponse *models.AuthRequestData, err error, ctx *gin.Context) {
	authReponse.Meta.Error = err.Error()
	authReponse.Meta.Message = "Error"
	ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
}

func GetFullname(userResponse *userResponse, userData *models.UserData) error {

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

	if err = json.Unmarshal([]byte(body), &userResponse); err != nil {
		return err
	}

	return err
}
