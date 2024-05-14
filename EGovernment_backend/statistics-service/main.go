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
	"statistics-service/handlers"
	"statistics-service/routes"
	"statistics-service/services"
)

var (
	server                      *gin.Engine
	ctx                         context.Context
	mongoClient                 *mongo.Client
	reportDelicTypeCollection   *mongo.Collection
	reportDelicTypeService      services.ReportDelicTypeService
	reportDelicTypeHandler      handlers.ReportDelicTypeHandler
	reportDelicTypeRouteHandler routes.ReportDelicTypeRouteHandler
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

	reportDelicTypeCollection = mongoClient.Database("Statistics").Collection("reportDelicType")

	reportDelicTypeService = services.NewReportDelicTypeImpl(reportDelicTypeCollection, ctx)
	reportDelicTypeHandler = handlers.NewReportDelicTypeHandler(reportDelicTypeService, reportDelicTypeCollection)
	reportDelicTypeRouteHandler = routes.NewReportDelicTypeRouteHandler(reportDelicTypeHandler, reportDelicTypeService)

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

	reportDelicTypeRouteHandler.Route(router)

	err := server.Run(":8082")
	if err != nil {
		fmt.Println(err)
		return
	}
}
