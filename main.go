package main

import (
	"context"
	"fmt"

	accountControllers "gypsylandFarming/app/controllers/accounts"
	authControllers "gypsylandFarming/app/controllers/auth"
	fileControllers "gypsylandFarming/app/controllers/files"
	otherControllers "gypsylandFarming/app/controllers/other"
	teamControllers "gypsylandFarming/app/controllers/teams"

	accountService "gypsylandFarming/app/services/accounts"
	authenticationService "gypsylandFarming/app/services/auth"
	filesService "gypsylandFarming/app/services/files"
	otherServices "gypsylandFarming/app/services/other"
	teamsService "gypsylandFarming/app/services/teams"

	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server      *gin.Engine
	ctx         context.Context
	mongoClient *mongo.Client

	locationService    otherServices.LocationService
	locationController otherControllers.LocationController
	locationCollection *mongo.Collection

	teamService    teamsService.TeamService
	teamController teamControllers.TeamController
	teamCollection *mongo.Collection

	teamAccessService    teamsService.TeamAccessService
	teamAccessController teamControllers.TeamAccessController
	teamAccessCollection *mongo.Collection

	readAccountRequestService    accountService.ReadAccountRequestService
	writeAccountRequestService   accountService.WriteAccountRequestService
	accountRequestController     accountControllers.AccountRequestController
	accountRequestTaskCollection *mongo.Collection

	accountTypesService   otherServices.AccountTypesService
	accountTypeController otherControllers.AccountTypesController
	accountTypeCollection *mongo.Collection

	authService    authenticationService.AuthService
	jwtService     authenticationService.JWTService
	authController authControllers.AuthController

	fileService    filesService.FileService
	fileController fileControllers.FileController
)

func init() {

	ctx = context.TODO()

	mongoConnection := options.Client().ApplyURI("mongodb://localhost/27017")
	mongoClient, err := mongo.Connect(ctx, mongoConnection)

	if err != nil {
		log.Fatal(err)
	}

	err = mongoClient.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mongo connection has been established")

	locationCollection = mongoClient.Database("gypsyland").Collection("locations")
	locationService = otherServices.NewLocationService(locationCollection, ctx)
	locationController = otherControllers.NewLocationController(locationService)

	teamAccessCollection = mongoClient.Database("gypsyland").Collection("teamAccess")
	teamAccessService = teamsService.NewTeamAccessService(teamAccessCollection, ctx)
	teamAccessController = teamControllers.NewTeamAccessController(teamAccessService)

	teamCollection = mongoClient.Database("gypsyland").Collection("teams")
	teamService = teamsService.NewTeamsService(teamCollection, ctx)
	teamController = teamControllers.NewTeamController(teamService, teamAccessService)

	accountRequestTaskCollection = mongoClient.Database("gypsyland").Collection("accountRequestTasks")

	accountTypeCollection = mongoClient.Database("gypsyland").Collection("accountTypes")
	accountTypesService = otherServices.NewAccountTypesService(accountTypeCollection, ctx)
	accountTypeController = otherControllers.NewAccountTypesController(accountTypesService)

	readAccountRequestService = accountService.NewReadAccountRequestService(accountRequestTaskCollection, ctx)
	writeAccountRequestService = accountService.NewWriteAccountRequestService(accountRequestTaskCollection, ctx)
	accountRequestController = accountControllers.NewAccountRequestTaskController(readAccountRequestService, writeAccountRequestService, teamService, teamAccessService, locationService, accountTypesService, fileService)

	authService = authenticationService.NewAuthService(ctx)
	jwtService = authenticationService.NewJWTService()
	authController = authControllers.NewAuthController(jwtService, authService, teamService)

	fileService = filesService.NewFileService(accountRequestTaskCollection, ctx)
	fileController = fileControllers.NewFileController(fileService)

	server = gin.Default()

	server.Use(cors.New(NewCORS()))

}

func NewCORS() cors.Config {
	config := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "OPTIIN", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}

	return config
}

func main() {

	defer mongoClient.Disconnect(ctx)

	basepath := server.Group("/v1")

	locationController.RegisterUserRoutes(basepath)

	accountRequestController.RegisterUserRoutes(basepath)
	accountTypeController.RegisterUserRoutes(basepath)
	authController.RegisterUserRoutes(basepath)
	teamController.RegisterUserRoutes(basepath)
	teamAccessController.RegisterUserRoutes(basepath)
	fileController.RegisterUserRoutes(basepath)

	log.Fatal(server.Run(":9090"))

}
