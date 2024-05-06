package main

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"net/http"
	//"vehicles-service/handlers"
	//"vehicles-service/routes"
	//"vehicles-service/services"
)

var (
	server            *gin.Engine
	ctx               context.Context
	mongoClient       *mongo.Client
	vehicleCollection *mongo.Collection
	//vehicleService      services.VehicleService
	//vehicleHandler      handlers.VehicleHandler
	//vehicleRouteHandler routes.VehicleRouteHandler
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

	vehicleCollection = mongoClient.Database("EGovernment").Collection("vehicle")

	//vehicleService = services.NewVehicleServiceImpl(vehicleCollection, ctx)
	//vehicleHandler = handlers.NewVehicleHandler(vehicleService, vehicleCollection)
	//vehicleRouteHandler = routes.NewVehicleRouteHandler(vehicleHandler, vehicleService)

	server = gin.Default()

}

func main() {
	defer func(mongoClient *mongo.Client, ctx context.Context) {
		err := mongoClient.Disconnect(ctx)
		if err != nil {

		}
	}(mongoClient, ctx)

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"https://localhost:4200"}
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthChecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Message"})
	})

	//statisticsRouteHandler.StatisticsRoute(router)

	err := server.Run(":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
}
