package routes

import (
	"statistics-service/handlers"
	"statistics-service/services"

	"github.com/gin-gonic/gin"
)

type ReportRegisteredVehiclesRouteHandler struct {
	handler handlers.ReportRegisteredVehiclesHandler
	service services.ReportRegisteredVehiclesService
}

func NewReportRegisteredVehiclesRouteHandler(handler handlers.ReportRegisteredVehiclesHandler, service services.ReportRegisteredVehiclesService) ReportRegisteredVehiclesRouteHandler {
	return ReportRegisteredVehiclesRouteHandler{handler, service}
}

func (rc *ReportRegisteredVehiclesRouteHandler) Route(rg *gin.RouterGroup) {
	router := rg.Group("/registeredVehiclesReport")
	router.POST("/create/category/:category/year/:year", rc.handler.CreateReport)
	router.GET("/all", rc.handler.GetAll)
	router.GET("/get/:id", rc.handler.GetByID)
	router.GET("/get/category/:category", rc.handler.GetAllByCategory)
	router.GET("/get/category/:category/year/:year", rc.handler.GetAllByCategoryAndYear)
}
