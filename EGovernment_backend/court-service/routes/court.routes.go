package routes

import (
	"court-service/handlers"
	"court-service/services"
	"github.com/gin-gonic/gin"
)

type CourtRouteHandler struct {
	handler handlers.CourtHandler
	service services.CourtService
}

func NewCourtRouteHandler(handler handlers.CourtHandler, service services.CourtService) CourtRouteHandler {
	return CourtRouteHandler{handler, service}
}

func (rc *CourtRouteHandler) CourtRoute(rg *gin.RouterGroup) {
	router := rg.Group("/court")
	router.POST("/courts", rc.handler.CreateCourt)
	router.GET("/courts/:id", rc.handler.GetCourtByID)
}
