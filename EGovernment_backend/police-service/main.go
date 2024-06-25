package main

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
	"os"
	"police-service/handlers"
	"police-service/routes"
	"police-service/services"
	"police-service/storage"
)

var (
	server                  *gin.Engine
	ctx                     context.Context
	mongoClient             *mongo.Client
	delictCollection        *mongo.Collection
	reportCollection        *mongo.Collection
	carAccidentCollection   *mongo.Collection
	delictService           services.DelictService
	reportService           services.ReportService
	carAccidentService      services.CarAccidentService
	delictHandler           handlers.DelictHandler
	reportHandler           handlers.ReportHandler
	carAccidentHandler      handlers.CarAccidentHandler
	delictRouteHandler      routes.DelictRouteHandler
	reportRouteHandler      routes.ReportRouteHandler
	carAccidentRouteHandler routes.CarAccidentRouteHandler
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

	delictCollection = mongoClient.Database("Police").Collection("delict")
	//log.Println("Delict Collection:", delictCollection)
	reportCollection = mongoClient.Database("Police").Collection("report")
	//log.Println("Report Collection:", reportCollection)
	carAccidentCollection = mongoClient.Database("Police").Collection("carAccident")

	reportService = services.NewReportServiceImpl(reportCollection, ctx)
	reportHandler = handlers.NewReportHandler(reportService, reportCollection)
	reportRouteHandler = routes.NewReportRouteHandler(reportHandler, reportService)

	//Initialize the logger we are going to use, with prefix and datetime for every log
	logger := log.New(os.Stdout, "[storage-api] ", log.LstdFlags)
	storageLogger := log.New(os.Stdout, "[FileStorage] ", log.Ldate|log.Ltime|log.Lshortfile)

	//// NoSQL: Initialize File Storage store
	fileStorage, err := storage.New(storageLogger)
	if err != nil {
		logger.Fatalf("Error initializing FileStorage: %v", err)
	}

	// Close connection to HDFS on shutdown
	defer fileStorage.Close()

	//// Create directory tree on HDFS
	_ = fileStorage.CreateDirectories()

	delictService = services.NewDelictServiceImpl(delictCollection, ctx)
	delictHandler = handlers.NewDelictHandler(delictService, delictCollection, reportService, logger, fileStorage)
	delictRouteHandler = routes.NewDelictRouteHandler(delictHandler, delictService, reportService)

	carAccidentService = services.NewCarAccidentServiceImpl(carAccidentCollection, ctx)
	carAccidentHandler = handlers.NewCarAccidentHandler(carAccidentService, carAccidentCollection)
	carAccidentRouteHandler = routes.NewCarAccidentRouteHandler(carAccidentHandler, carAccidentService)

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

	delictRouteHandler.DelictRoute(router)
	reportRouteHandler.ReportRoute(router)
	carAccidentRouteHandler.CarAccidentRoute(router)

	err := server.Run(":8084")
	if err != nil {
		fmt.Println(err)
		return
	}
}
