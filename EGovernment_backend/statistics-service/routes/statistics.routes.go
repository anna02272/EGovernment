package routes

import (
	"github.com/gin-gonic/gin"
	"statistics-service/handlers"
	"statistics-service/services"
)

type StatisticsRouteHandler struct {
	handler handlers.StatisticsHandler
	service services.StatisticsService
}

func NewStatisticsRouteHandler(handler handlers.StatisticsHandler, service services.StatisticsService) StatisticsRouteHandler {
	return StatisticsRouteHandler{handler, service}
}

func (rc *StatisticsRouteHandler) StatisticsRoute(rg *gin.RouterGroup) {
	_ = rg.Group("/statistics")

}
