package routes

import (
	"statistics-service/handlers"
	"statistics-service/services"

	"github.com/gin-gonic/gin"
)

type ReportCarAccidentDegreeRouteHandler struct {
	handler handlers.ReportCarAccidentDegreeHandler
	service services.ReportCarAccidentDegreeService
}

func NewReportCarAccidentDegreeRouteHandler(handler handlers.ReportCarAccidentDegreeHandler, service services.ReportCarAccidentDegreeService) ReportCarAccidentDegreeRouteHandler {
	return ReportCarAccidentDegreeRouteHandler{handler, service}
}

func (rc *ReportCarAccidentDegreeRouteHandler) Route(rg *gin.RouterGroup) {
	router := rg.Group("/carAccidentDegreeReport")
	router.POST("/create/degree/:degree/year/:year", rc.handler.CreateReport)
	router.GET("/all", rc.handler.GetAll)
	router.GET("/get/:id", rc.handler.GetByID)
	router.GET("/get/degree/:degree", rc.handler.GetAllByCarAccidentDegree)
	router.GET("/get/degree/:degree/year/:year", rc.handler.GetAllByCarAccidentDegreeAndYear)
}
