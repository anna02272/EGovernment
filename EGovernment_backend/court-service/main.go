package main

import (
	"context"
	"court-service/handlers"
	"court-service/routes"
	"court-service/services"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"net/http"
)

var (
	server              *gin.Engine
	ctx                 context.Context
	mongoClient         *mongo.Client
	courtCollection     *mongo.Collection
	subjectCollection   *mongo.Collection
	citizenCollection   *mongo.Collection
	courtService        services.CourtService
	courtHandler        handlers.CourtHandler
	courtRouteHandler   routes.CourtRouteHandler
	citizenHandler      handlers.CitizenHandler
	citizenService      services.CitizenService
	citizenRouteHandler routes.CitizenRouteHandler
	subjectHandler      handlers.SubjectHandler
	subjectService      services.SubjectService
	subjectRouteHandler routes.SubjectRouteHandler
)

func init() {
	ctx = context.TODO()
	mongoConn := options.Client().ApplyURI("mongodb://root:root@mongo:27017")
	mongoClient, err := mongo.Connect(ctx, mongoConn)

	if err != nil {
		panic(err)
	}

	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("MongoDB successfully connected...")

	courtCollection = mongoClient.Database("EGovernment").Collection("court")

	courtService = services.NewCourtServiceImpl(courtCollection, ctx)
	courtHandler = handlers.NewCourtHandler(courtService, courtCollection)
	courtRouteHandler = routes.NewCourtRouteHandler(courtHandler, courtService)

	citizenCollection = mongoClient.Database("EGovernment").Collection("citizen")

	citizenService = services.NewCitizenServiceImpl(citizenCollection, ctx)
	citizenHandler = handlers.NewCitizenHandler(citizenService, citizenCollection)
	citizenRouteHandler = routes.NewCitizenRouteHandler(citizenHandler, citizenService)

	subjectCollection = mongoClient.Database("EGovernment").Collection("subject")

	subjectService = services.NewSubjectServiceImpl(subjectCollection, ctx)
	subjectHandler = handlers.NewSubjectHandler(subjectService, subjectCollection)
	subjectRouteHandler = routes.NewSubjectRouteHandler(subjectHandler, subjectService)

	server = gin.Default()

}

func main() {
	defer func(mongoClient *mongo.Client, ctx context.Context) {
		err := mongoClient.Disconnect(ctx)
		if err != nil {

		}
	}(mongoClient, ctx)

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:4200"}
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthChecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Message"})
	})

	courtRouteHandler.CourtRoute(router)
	citizenRouteHandler.CitizenRoute(router)
	subjectRouteHandler.SubjectRoutes(router)

	err := server.Run(":8083")
	if err != nil {
		fmt.Println(err)
		return
	}
}
