package main

import (
	"context"
	"fmt"
	"gypsyland_farming/app/controllers"
	services "gypsyland_farming/app/services"
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

	locationService    services.LocationService
	locationController controllers.LocationController
	locationCollection *mongo.Collection

	positionService    services.PositionService
	positionController controllers.PositionController
	positionCollection *mongo.Collection

	teamService    services.TeamService
	teamController controllers.TeamController
	teamCollection *mongo.Collection

	teamAccessService    services.TeamAccessService
	teamAccessController controllers.TeamAccessController
	teamAccessCollection *mongo.Collection

	readAccountRequestService    services.ReadAccountRequestService
	writeAccountRequestService   services.WriteAccountRequestService
	accountRequestController     controllers.AccountRequestController
	accountRequestCollection     *mongo.Collection
	accountRequestTaskCollection *mongo.Collection

	accountTypesService   services.AccountTypesService
	accountTypeController controllers.AccountTypesController
	accountTypeCollection *mongo.Collection

	authService    services.AuthService
	jwtService     services.JWTService
	authController controllers.AuthController

	fileService    services.FileService
	fileController controllers.FileController
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
	locationService = services.NewLocationService(locationCollection, ctx)
	locationController = controllers.NewLocationController(locationService)

	positionCollection = mongoClient.Database("gypsyland").Collection("positions")
	positionService = services.NewPositionService(positionCollection, ctx)
	positionController = controllers.NewPositionController(positionService)

	teamAccessCollection = mongoClient.Database("gypsyland").Collection("teamAccess")
	teamAccessService = services.NewTeamAccessService(teamAccessCollection, ctx)
	teamAccessController = controllers.NewTeamAccessController(teamAccessService)

	teamCollection = mongoClient.Database("gypsyland").Collection("teams")
	teamService = services.NewTeamsService(teamCollection, ctx)
	teamController = controllers.NewTeamController(teamService, teamAccessService)

	accountRequestCollection = mongoClient.Database("gypsyland").Collection("accountRequests")
	accountRequestTaskCollection = mongoClient.Database("gypsyland").Collection("accountRequestTasks")

	accountTypeCollection = mongoClient.Database("gypsyland").Collection("accountTypes")
	accountTypesService = services.NewAccountTypesService(accountTypeCollection, ctx)
	accountTypeController = controllers.NewAccountTypesController(accountTypesService)

	readAccountRequestService = services.NewReadAccountRequestService(accountRequestCollection, accountRequestTaskCollection, ctx)
	writeAccountRequestService = services.NewWriteAccountRequestService(accountRequestCollection, accountRequestTaskCollection, ctx)
	accountRequestController = controllers.NewAccountRequestTaskController(readAccountRequestService, writeAccountRequestService, teamService, teamAccessService, locationService, accountTypesService, fileService)

	authService = services.NewAuthService(ctx)
	jwtService = services.NewJWTService()
	authController = controllers.NewAuthController(jwtService, authService, teamService)

	fileService = services.NewFileService(accountRequestTaskCollection, ctx)
	fileController = controllers.NewFileController(fileService)

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
	positionController.RegisterUserRoutes(basepath)
	accountRequestController.RegisterUserRoutes(basepath)
	accountTypeController.RegisterUserRoutes(basepath)
	authController.RegisterUserRoutes(basepath)
	teamController.RegisterUserRoutes(basepath)
	teamAccessController.RegisterUserRoutes(basepath)
	fileController.RegisterUserRoutes(basepath)

	log.Fatal(server.Run(":9090"))

}
