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
	server      *gin.Engine
	ctx         context.Context
	mongoClient *mongo.Client

	reportDelicTypeCollection   *mongo.Collection
	reportDelicTypeService      services.ReportDelicTypeService
	reportDelicTypeHandler      handlers.ReportDelicTypeHandler
	reportDelicTypeRouteHandler routes.ReportDelicTypeRouteHandler

	reportCarAccidentTypeCollection   *mongo.Collection
	reportCarAccidentTypeService      services.ReportCarAccidentTypeService
	reportCarAccidentTypeHandler      handlers.ReportCarAccidentTypeHandler
	reportCarAccidentTypeRouteHandler routes.ReportCarAccidentTypeRouteHandler

	reportCarAccidentDegreeCollection   *mongo.Collection
	reportCarAccidentDegreeService      services.ReportCarAccidentDegreeService
	reportCarAccidentDegreeHandler      handlers.ReportCarAccidentDegreeHandler
	reportCarAccidentDegreeRouteHandler routes.ReportCarAccidentDegreeRouteHandler

	reportRegisteredVehiclesCollection   *mongo.Collection
	reportRegisteredVehiclesService      services.ReportRegisteredVehiclesService
	reportRegisteredVehiclesHandler      handlers.ReportRegisteredVehiclesHandler
	reportRegisteredVehiclesRouteHandler routes.ReportRegisteredVehiclesRouteHandler

	responseCollection   *mongo.Collection
	responseService      services.ResponseService
	responseHandler      handlers.ResponseHandler
	responseRouteHandler routes.ResponseRouteHandler

	requestCollection   *mongo.Collection
	requestService      services.RequestService
	requestHandler      handlers.RequestHandler
	requestRouteHandler routes.RequestRouteHandler
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
	reportCarAccidentTypeCollection = mongoClient.Database("Statistics").Collection("reportCarAccidentType")
	reportCarAccidentDegreeCollection = mongoClient.Database("Statistics").Collection("reportCarAccidentDegree")
	reportRegisteredVehiclesCollection = mongoClient.Database("Statistics").Collection("reportRegisteredVehicles")
	responseCollection = mongoClient.Database("Statistics").Collection("response")
	requestCollection = mongoClient.Database("Statistics").Collection("request")

	reportDelicTypeService = services.NewReportDelicTypeImpl(reportDelicTypeCollection, ctx)
	reportDelicTypeHandler = handlers.NewReportDelicTypeHandler(reportDelicTypeService, reportDelicTypeCollection)
	reportDelicTypeRouteHandler = routes.NewReportDelicTypeRouteHandler(reportDelicTypeHandler, reportDelicTypeService)

	reportCarAccidentTypeService = services.NewReportCarAccidentTypeImpl(reportCarAccidentTypeCollection, ctx)
	reportCarAccidentTypeHandler = handlers.NewReportCarAccidentTypeHandler(reportCarAccidentTypeService, reportCarAccidentTypeCollection)
	reportCarAccidentTypeRouteHandler = routes.NewReportCarAccidentTypeRouteHandler(reportCarAccidentTypeHandler, reportCarAccidentTypeService)

	reportCarAccidentDegreeService = services.NewReportCarAccidentDegreeImpl(reportCarAccidentDegreeCollection, ctx)
	reportCarAccidentDegreeHandler = handlers.NewReportCarAccidentDegreeHandler(reportCarAccidentDegreeService, reportCarAccidentDegreeCollection)
	reportCarAccidentDegreeRouteHandler = routes.NewReportCarAccidentDegreeRouteHandler(reportCarAccidentDegreeHandler, reportCarAccidentDegreeService)

	reportRegisteredVehiclesService = services.NewReportRegisteredVehiclesServiceImpl(reportRegisteredVehiclesCollection, ctx)
	reportRegisteredVehiclesHandler = handlers.NewReportRegisteredVehiclesHandler(reportRegisteredVehiclesService, reportRegisteredVehiclesCollection)
	reportRegisteredVehiclesRouteHandler = routes.NewReportRegisteredVehiclesRouteHandler(reportRegisteredVehiclesHandler, reportRegisteredVehiclesService)

	responseService = services.NewResponseServiceImpl(responseCollection, ctx)
	responseHandler = handlers.NewResponseHandler(responseService, responseCollection)
	responseRouteHandler = routes.NewResponseRouteHandler(responseHandler, responseService)

	requestService = services.NewRequestServiceImpl(requestCollection, ctx)
	requestHandler = handlers.NewRequestHandler(requestService, requestCollection)
	requestRouteHandler = routes.NewRequestRouteHandler(requestHandler, requestService)

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
	reportCarAccidentTypeRouteHandler.Route(router)
	reportCarAccidentDegreeRouteHandler.Route(router)
	reportRegisteredVehiclesRouteHandler.Route(router)
	responseRouteHandler.Route(router)
	requestRouteHandler.Route(router)

	err := server.Run(":8082")
	if err != nil {
		fmt.Println(err)
		return
	}
}
