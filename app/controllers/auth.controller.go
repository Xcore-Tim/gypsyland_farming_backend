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

func NewAuthController(authService services.AuthService, jwtService services.JWTService, teamService services.TeamService) AuthController {
	return AuthController{
		AuthService: authService,
		JWTService:  jwtService,
		TeamService: teamService,
	}
}

func (auc AuthController) Login(ctx *gin.Context) {

	var authResponse models.AuthResponse
	client := &http.Client{}

	endpoint := "/v1/accounts/auth"
	urlPath := models.Basepath + endpoint

	username := ctx.Request.PostFormValue("username")
	password := ctx.Request.PostFormValue("password")

	requestBody, _ := json.Marshal(map[string]string{
		"email":    username,
		"password": password,
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

	var userIdentity models.UserIdentity

	if err = json.Unmarshal([]byte(body), &userIdentity); err != nil {
		DenyAuthentication(&authResponse, err, ctx)
		return
	}

	var userResponse userResponse

	if err := GetFullname(&userResponse, &userIdentity); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authResponse.Code = http.StatusOK
	authResponse.Data.Token.AccessToken = userIdentity.Token
	authResponse.Data.Token.RequestToken = userIdentity.Token
	authResponse.Data.FullName = userResponse.Username
	authResponse.Data.UserID = userIdentity.UserID
	authResponse.Data.RoleID = userIdentity.RoleID
	authResponse.Data.TeamID = userIdentity.TeamID

	if userIdentity.RoleID == 2 {

		if _, err := auc.TeamService.GetTeamByNum(userIdentity.TeamID); err != nil {

			teamlead := models.Employee{
				ID:       userIdentity.UserID,
				Name:     userResponse.Username,
				Position: userIdentity.RoleID,
			}

			team := models.Team{
				ID:       userIdentity.TeamID,
				Number:   userIdentity.TeamID,
				TeamLead: teamlead,
			}

			auc.TeamService.CreateTeam(&team)

		}

	}

	ctx.JSON(http.StatusOK, authResponse)
}

func (auc AuthController) CreateUser(ctx *gin.Context) {

	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	err := auc.AuthService.CreateUser(&user)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)

}

func (auc AuthController) GetAll(ctx *gin.Context) {

	users, err := auc.AuthService.GetAll()

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)

}

func (auc AuthController) RegisterUserRoutes(rg *gin.RouterGroup) {

	authRequestGroup := rg.Group("/auth")

	authRequestGroup.POST("", auc.Login)

	usersGroup := authRequestGroup.Group("/users")
	usersGroup.POST("/create", auc.CreateUser)
	usersGroup.POST("/getall", auc.GetAll)

}

func DenyAuthentication(authReponse *models.AuthResponse, err error, ctx *gin.Context) {

	authReponse.Code = http.StatusBadRequest
	authReponse.Data.Meta.Error = err.Error()
	authReponse.Data.Meta.Message = "Error"
	ctx.JSON(http.StatusBadRequest, gin.H{"error": err})

}
func GetFullname(userResponse *userResponse, userIdentity *models.UserIdentity) error {

	endpoint := "/v1/Identity/users/"
	urlPath := models.Basepath + endpoint + strconv.Itoa(userIdentity.UserID)

	bearer := "BEARER " + userIdentity.Token

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
