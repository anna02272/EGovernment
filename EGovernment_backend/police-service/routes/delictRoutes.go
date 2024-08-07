package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"police-service/domain"
	"police-service/handlers"
	"police-service/services"
)

type DelictRouteHandler struct {
	handler       handlers.DelictHandler
	service       services.DelictService
	reportService services.ReportService
}

func NewDelictRouteHandler(handler handlers.DelictHandler, service services.DelictService, reportService services.ReportService) DelictRouteHandler {
	return DelictRouteHandler{handler, service, reportService}
}

func (d *DelictRouteHandler) DelictRoute(rg *gin.RouterGroup) {
	router := rg.Group("/delict")
	router.POST("/createDelict", MiddlewareDelictDeserialization, d.handler.CreateDelict)
	router.GET("/all", d.handler.GetAllDelicts)
	router.GET("/getPolicemanDelicts", d.handler.GetDelictsByPolicemanID)
	router.GET("/getDriverDelicts", d.handler.GetDelictsByDriver)
	router.GET("/get/:id", d.handler.GetDelictByID)
	router.GET("/getDriver/:driverId", d.handler.CheckDriverAlcoholDelicts)
	router.GET("/get/delictType/:delictType", d.handler.GetAllDelictsByDelictType)
	router.GET("/get/delictType/:delictType/year/:year", d.handler.GetAllDelictsByDelictTypeAndYear)
	router.PATCH("/pay/:id", d.handler.PayFine)
	router.GET("/getImagesUrls/:folderName", d.handler.GetImageURLS)
	router.GET("/getImages/:folderName/:imageName", d.handler.GetImageContent)
	router.POST("/upload/:folderName", d.handler.UploadImages)
	router.GET("/test/pdf", d.handler.GenerateTestPDF)
	router.GET("/getPdf/:id", d.handler.ServeDelictPDF)
}

func MiddlewareDelictDeserialization(c *gin.Context) {
	var delict domain.DelictCreate

	if err := c.ShouldBindJSON(&delict); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to decode JSON"})
		c.Abort()
		return
	}

	c.Set("delict", delict)
	c.Next()
}
