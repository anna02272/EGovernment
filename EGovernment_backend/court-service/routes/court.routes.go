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
	_ = rg.Group("/court")

}
