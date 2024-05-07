package routes

import (
	"auth-service/handlers"
	"auth-service/services"
	"github.com/gin-gonic/gin"
)

type AuthRouteHandler struct {
	authHandler handlers.AuthHandler
	authService services.AuthService
}

func NewAuthRouteHandler(authHandler handlers.AuthHandler, authService services.AuthService) AuthRouteHandler {
	return AuthRouteHandler{authHandler, authService}
}

func (rc *AuthRouteHandler) AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("/auth")
	router.Use(handlers.ExtractTraceInfoMiddleware())

	router.POST("/login", rc.authHandler.Login)
	router.POST("/register", rc.authHandler.Registration)

}
