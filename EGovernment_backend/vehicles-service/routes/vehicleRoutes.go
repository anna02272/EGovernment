package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vehicles-service/domain"
	"vehicles-service/handlers"
	"vehicles-service/services"
)

type VehicleRouteHandler struct {
	handler       handlers.VehicleHandler
	service       services.VehicleService
	driverService services.VehicleDriverService
}

func NewVehicleRouteHandler(handler handlers.VehicleHandler,
	service services.VehicleService, driverService services.VehicleDriverService) VehicleRouteHandler {
	return VehicleRouteHandler{handler, service, driverService}
}

func (vr *VehicleRouteHandler) VehicleRoute(rg *gin.RouterGroup) {
	router := rg.Group("/vehicle")
	router.POST("/createVehicle", MiddlewareVehicleDeserialization, vr.handler.CreateVehicle)
	router.GET("/all", vr.handler.GetAllVehicles)
	router.GET("/all/registered", vr.handler.GetAllRegisteredVehicles)
	router.GET("/get/category/:category/year/:year", vr.handler.GetAllVehiclesByCategoryAndYear)
	router.GET("/get/:id", vr.handler.GetVehicleByID)
	router.GET("/registeredVehicles/pdf", vr.handler.GenerateAndServeVehiclesReportPDF)
	router.GET("/registeredVehicles/category/:category/pdf", vr.handler.GenerateAndServeVehiclesByCategoryReportPDF)
	router.GET("/count/category/:category", vr.handler.GetNumberOfRegisteredVehiclesByCategory)
	router.GET("/all/registered/:category", vr.handler.GetAllRegisteredVehiclesByCategory)

}

func MiddlewareVehicleDeserialization(c *gin.Context) {
	var vehicle domain.VehicleCreate

	if err := c.ShouldBindJSON(&vehicle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to decode JSON"})
		c.Abort()
		return
	}

	c.Set("vehicle", vehicle)
	c.Next()
}
