package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"police-service/domain"
	"police-service/handlers"
	"police-service/services"
)

type ReportRouteHandler struct {
	handler handlers.ReportHandler
	service services.ReportService
}

func NewReportRouteHandler(handler handlers.ReportHandler, service services.ReportService) ReportRouteHandler {
	return ReportRouteHandler{handler, service}
}

func (d *ReportRouteHandler) ReportRoute(rg *gin.RouterGroup) {
	router := rg.Group("/report")
	router.POST("/createReport", MiddlewareReportDeserialization, d.handler.CreateReport)
	router.GET("/all", d.handler.GetAllReports)
	router.GET("/get/:id", d.handler.GetReportByID)
	router.GET("/get/delictType/:delictType", d.handler.GetAllReportssByDelictType)
}

func MiddlewareReportDeserialization(c *gin.Context) {
	var report domain.ReportCreate

	if err := c.ShouldBindJSON(&report); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to decode JSON"})
		c.Abort()
		return
	}

	c.Set("report", report)
	c.Next()
}
