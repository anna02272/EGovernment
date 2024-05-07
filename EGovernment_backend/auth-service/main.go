package main

import (
	"auth-service/handlers"
	"auth-service/routes"
	"auth-service/services"
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"net/http"
)

var (
	server      *gin.Engine
	ctx         context.Context
	mongoClient *mongo.Client

	userService      services.UserService
	UserHandler      handlers.UserHandler
	UserRouteHandler routes.UserRouteHandler

	authCollection   *mongo.Collection
	authService      services.AuthService
	AuthHandler      handlers.AuthHandler
	AuthRouteHandler routes.AuthRouteHandler
)

func init() {
	ctx = context.TODO()
	mongoConn := options.Client().ApplyURI("mongodb://root:root@mongo:27017")
	mongoClient, err := mongo.Connect(ctx, mongoConn)

	if err != nil {
		fmt.Printf("Error connecting to MongoDB: %v\n", err)
		panic(err)
	}

	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		fmt.Printf("Error disconnecting from MongoDB: %v\n", err)
		panic(err)
	}

	fmt.Println("MongoDB successfully connected...")

	authCollection = mongoClient.Database("EGovernment").Collection("auth")
	userService = services.NewUserServiceImpl(authCollection, ctx)
	authService = services.NewAuthService(authCollection, ctx, userService)
	AuthHandler = handlers.NewAuthHandler(authService, userService, authCollection)
	AuthRouteHandler = routes.NewAuthRouteHandler(AuthHandler, authService)
	UserHandler = handlers.NewUserHandler(userService)
	UserRouteHandler = routes.NewRouteUserHandler(UserHandler)

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

	AuthRouteHandler.AuthRoute(router)
	UserRouteHandler.UserRoute(router)

	err := server.Run(":8085")
	if err != nil {
		fmt.Println(err)
		return
	}

}
