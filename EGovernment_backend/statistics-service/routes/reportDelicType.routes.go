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
	router := rg.Group("/statistics")
	router.POST("/create/delictReport/:delictType", rc.handler.CreateDelictsReport)
}
