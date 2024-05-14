package routes

import (
	"statistics-service/handlers"
	"statistics-service/services"

	"github.com/gin-gonic/gin"
)

type ReportDelicTypeRouteHandler struct {
	handler handlers.ReportDelicTypeHandler
	service services.ReportDelicTypeService
}

func NewReportDelicTypeRouteHandler(handler handlers.ReportDelicTypeHandler, service services.ReportDelicTypeService) ReportDelicTypeRouteHandler {
	return ReportDelicTypeRouteHandler{handler, service}
}

func (rc *ReportDelicTypeRouteHandler) Route(rg *gin.RouterGroup) {
	router := rg.Group("/delictReport")
	router.POST("/create/delictType/:delictType/year/:year", rc.handler.CreateDelictsReport)
	router.GET("/all", rc.handler.GetAll)
	router.GET("/get/:id", rc.handler.GetByID)
	router.GET("/get/delictType/:delictType", rc.handler.GetAllByDelictType)
	router.GET("/get/delictType/:delictType/year/:year", rc.handler.GetAllByDelictTypeAndYear)
}
