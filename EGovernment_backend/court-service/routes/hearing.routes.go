package routes

import (
	"court-service/handlers"
	"court-service/services"
	"github.com/gin-gonic/gin"
)

type HearingRouteHandler struct {
	handler handlers.HearingHandler
	service services.HearingService
}

func NewHearingRouteHandler(handler handlers.HearingHandler, service services.HearingService) HearingRouteHandler {
	return HearingRouteHandler{handler, service}
}

func (rc *HearingRouteHandler) HearingRoute(rg *gin.RouterGroup) {
	router := rg.Group("/hearing")
	router.POST("/create", rc.handler.CreateHearing)
	router.GET("/hearings/:id", rc.handler.GetHearingByID)
	router.GET("/getByIdJudge", rc.handler.GetJudgeHearings)

}
