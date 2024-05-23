package routes

import (
	"statistics-service/handlers"
	"statistics-service/services"

	"github.com/gin-gonic/gin"
)

type ReportCarAccidentTypeRouteHandler struct {
	handler handlers.ReportCarAccidentTypeHandler
	service services.ReportCarAccidentTypeService
}

func NewReportCarAccidentTypeRouteHandler(handler handlers.ReportCarAccidentTypeHandler, service services.ReportCarAccidentTypeService) ReportCarAccidentTypeRouteHandler {
	return ReportCarAccidentTypeRouteHandler{handler, service}
}

func (rc *ReportCarAccidentTypeRouteHandler) Route(rg *gin.RouterGroup) {
	router := rg.Group("/carAccidentTypeReport")
	router.POST("/create/carAccidentType/:carAccidentType/year/:year", rc.handler.CreateReport)
	router.GET("/all", rc.handler.GetAll)
	router.GET("/get/:id", rc.handler.GetByID)
	router.GET("/get/carAccidentType/:carAccidentType", rc.handler.GetAllByCarAccidentType)
	router.GET("/get/carAccidentType/:carAccidentType/year/:year", rc.handler.GetAllByCarAccidentTypeAndYear)
}
