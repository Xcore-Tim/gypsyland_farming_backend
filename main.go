package main

import (
	"context"
	"fmt"

	accountTypesControllers "gypsylandFarming/app/controllers/accountTypes"
	accountControllers "gypsylandFarming/app/controllers/accounts"
	authControllers "gypsylandFarming/app/controllers/auth"
	currencyControllers "gypsylandFarming/app/controllers/currency"
	fileControllers "gypsylandFarming/app/controllers/files"
	locationsControllers "gypsylandFarming/app/controllers/locations"
	teamControllers "gypsylandFarming/app/controllers/teams"

	accountTypesServices "gypsylandFarming/app/services/accountTypes"
	accountServices "gypsylandFarming/app/services/accounts"
	authenticationServices "gypsylandFarming/app/services/auth"
	currencyServices "gypsylandFarming/app/services/currency"
	filesServices "gypsylandFarming/app/services/files"
	locationsServices "gypsylandFarming/app/services/locations"
	teamsServices "gypsylandFarming/app/services/teams"

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

	locationService    locationsServices.LocationService
	locationController locationsControllers.LocationController
	locationCollection *mongo.Collection

	teamService    teamsServices.TeamService
	teamController teamControllers.TeamController
	teamCollection *mongo.Collection

	teamAccessService    teamsServices.TeamAccessService
	teamAccessController teamControllers.TeamAccessController
	teamAccessCollection *mongo.Collection

	readAccountRequestService    accountServices.ReadAccountRequestService
	writeAccountRequestService   accountServices.WriteAccountRequestService
	accountRequestController     accountControllers.AccountRequestController
	accountRequestTaskCollection *mongo.Collection

	accountTypesService   accountTypesServices.AccountTypesService
	accountTypeController accountTypesControllers.AccountTypesController
	accountTypeCollection *mongo.Collection

	authService    authenticationServices.AuthService
	jwtService     authenticationServices.JWTService
	authController authControllers.AuthController

	fileService    filesServices.FileService
	fileController fileControllers.FileController

	currencyService      currencyServices.CurrencyService
	currencyRatesService currencyServices.CurrencyRatesService
	currencyController   currencyControllers.CurrencyController
	currencyCollection   *mongo.Collection
)

func init() {

	ctx = context.TODO()

	connectionString := "mongodb://farming-mongodb:7wnnOjnZgpq4Ruprtqq5qXxsS7ZfCF8LxhHJYIZgzenmAJc3l1ZrFEsT5AuCYXjWtvGAJ6Fdfj0lACDbWPXUiw==@farming-mongodb.mongo.cosmos.azure.com:10255/?ssl=true&retrywrites=false&maxIdleTimeMS=120000&appName=@farming-mongodb@"
	mongoConnection := options.Client().ApplyURI(connectionString)
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
	locationService = locationsServices.NewLocationService(locationCollection, ctx)
	locationController = locationsControllers.NewLocationController(locationService)

	teamAccessCollection = mongoClient.Database("gypsyland").Collection("teamAccess")
	teamAccessService = teamsServices.NewTeamAccessService(teamAccessCollection, ctx)
	teamAccessController = teamControllers.NewTeamAccessController(teamAccessService)

	teamCollection = mongoClient.Database("gypsyland").Collection("teams")
	teamService = teamsServices.NewTeamsService(teamCollection, ctx)
	teamController = teamControllers.NewTeamController(teamService, teamAccessService)

	accountRequestTaskCollection = mongoClient.Database("gypsyland").Collection("accountRequestTasks")

	accountTypeCollection = mongoClient.Database("gypsyland").Collection("accountTypes")
	accountTypesService = accountTypesServices.NewAccountTypesService(accountTypeCollection, ctx)
	accountTypeController = accountTypesControllers.NewAccountTypesController(accountTypesService)

	currencyCollection = mongoClient.Database("gypsyland").Collection("currency")
	currencyService = currencyServices.NewCurrencyService(currencyCollection, ctx)
	currencyRatesService = currencyServices.NewCurrencyRatesService(currencyCollection, ctx)
	currencyController = currencyControllers.NewCurrencyController(currencyService, currencyRatesService)

	readAccountRequestService = accountServices.NewReadAccountRequestService(accountRequestTaskCollection, ctx)
	writeAccountRequestService = accountServices.NewWriteAccountRequestService(accountRequestTaskCollection, ctx)
	accountRequestController = accountControllers.NewAccountRequestTaskController(
		readAccountRequestService,
		writeAccountRequestService,
		teamService,
		teamAccessService,
		locationService,
		accountTypesService,
		fileService,
		currencyService,
		currencyRatesService,
	)

	authService = authenticationServices.NewAuthService(ctx)
	jwtService = authenticationServices.NewJWTService()
	authController = authControllers.NewAuthController(jwtService, authService, teamService)

	fileService = filesServices.NewFileService(accountRequestTaskCollection, ctx)
	fileController = fileControllers.NewFileController(fileService)

	server = gin.Default()

	server.Use(cors.New(NewCORS()))

}

func NewCORS() cors.Config {
	config := cors.Config{
		// AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "OPTIIN", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
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
	currencyController.RegisterUserRoutes(basepath)

	// port := os.Getenv("HTPP_PLATFORM_PORT")

	// if port == "" {
	// 	port = ":80"
	// }
	port := ":80"
	log.Fatal(server.Run(port))

}
