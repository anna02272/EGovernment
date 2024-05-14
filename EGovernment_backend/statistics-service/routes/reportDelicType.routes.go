package routes

import (
	"github.com/gin-gonic/gin"
	"statistics-service/handlers"
	"statistics-service/services"
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
	router.POST("/create/:delictType", rc.handler.CreateDelictsReport)
	router.GET("/all", rc.handler.GetAll)
	router.GET("/get/:id", rc.handler.GetByID)
	router.GET("/get/delictType/:delictType", rc.handler.GetAllByDelictType)
}
