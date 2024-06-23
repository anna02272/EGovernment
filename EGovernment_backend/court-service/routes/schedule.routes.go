package routes

import (
	"court-service/handlers"
	"court-service/services"
	"github.com/gin-gonic/gin"
)

type ScheduleRouteHandler struct {
	handler handlers.ScheduleHandler
	service services.ScheduleService
}

func NewScheduleRouteHandler(handler handlers.ScheduleHandler, service services.ScheduleService) ScheduleRouteHandler {
	return ScheduleRouteHandler{handler, service}
}

func (rc *ScheduleRouteHandler) ScheduleRoute(rg *gin.RouterGroup) {
	router := rg.Group("/schedule")
	router.POST("/create", rc.handler.CreateSchedule)
	router.GET("/schedules/:id", rc.handler.GetScheduleByID)
	router.GET("/getByHearing/:hearingId", rc.handler.GetScheduleByHearingID)

}
