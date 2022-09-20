package main

import (
	"context"
	"fmt"
	"gypsyland_farming/app/controllers"
	"gypsyland_farming/app/services"
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

	employeeService    services.EmployeeService
	employeeController controllers.EmployeeController
	employeeCollection *mongo.Collection

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

	accountTypeService    services.AccountTypesService
	accountTypeController controllers.AccountTypesController
	accountTypeCollection *mongo.Collection

	authService    services.AuthService
	jwtService     services.JWTService
	userCollection *mongo.Collection
	authController controllers.AuthController
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

	employeeCollection = mongoClient.Database("gypsyland").Collection("employees")
	employeeService = services.NewEmployeeService(employeeCollection, ctx)
	employeeController = controllers.NewEmployeeController(employeeService)

	teamCollection = mongoClient.Database("gypsyland").Collection("teams")
	teamService = services.NewTeamsService(teamCollection, ctx)
	teamController = controllers.NewTeamController(teamService)

	teamAccessCollection = mongoClient.Database("gypsyland").Collection("teamAccess")
	teamAccessService = services.NewTeamAccessService(teamAccessCollection, ctx)
	teamAccessController = controllers.NewTeamAccessController(teamAccessService)

	accountRequestCollection = mongoClient.Database("gypsyland").Collection("accountRequests")
	accountRequestTaskCollection = mongoClient.Database("gypsyland").Collection("accountRequestTasks")

	readAccountRequestService = services.NewReadAccountRequestService(accountRequestCollection, accountRequestTaskCollection, ctx)
	writeAccountRequestService = services.NewWriteAccountRequestService(accountRequestCollection, accountRequestTaskCollection, ctx)
	accountRequestController = controllers.NewAccountRequestTaskController(readAccountRequestService, writeAccountRequestService, teamService)

	accountTypeCollection = mongoClient.Database("gypsyland").Collection("accountTypes")
	accountTypeService = services.NewAccountTypesService(accountTypeCollection, ctx)
	accountTypeController = controllers.NewAccountTypesController(accountTypeService)

	userCollection = mongoClient.Database("gypsyland").Collection("users")
	authService = services.NewAuthService(userCollection, employeeCollection, ctx)
	jwtService = services.NewJWTService()
	authController = controllers.NewAuthController(authService, jwtService, teamService)

	server = gin.Default()
	server.Use(cors.Default())

}

// v1/location/create
func main() {

	defer mongoClient.Disconnect(ctx)

	basepath := server.Group("/v1")

	locationController.RegisterUserRoutes(basepath)
	positionController.RegisterUserRoutes(basepath)
	employeeController.RegisterUserRoutes(basepath)
	accountRequestController.RegisterUserRoutes(basepath)
	accountTypeController.RegisterUserRoutes(basepath)
	authController.RegisterUserRoutes(basepath)
	teamController.RegisterUserRoutes(basepath)
	teamAccessController.RegisterUserRoutes(basepath)

	log.Fatal(server.Run(":9090"))

}
