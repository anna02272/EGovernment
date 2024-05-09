package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vehicles-service/domain"
	"vehicles-service/handlers"
	"vehicles-service/services"
)

type VehicleDriverRouteHandler struct {
	handler handlers.VehicleDriverHandler
	service services.VehicleDriverService
}

func NewVehicleDriverRouteHandler(handler handlers.VehicleDriverHandler, service services.VehicleDriverService) VehicleDriverRouteHandler {
	return VehicleDriverRouteHandler{handler, service}
}

func (vr *VehicleDriverRouteHandler) VehicleDriverRoute(rg *gin.RouterGroup) {
	router := rg.Group("/driver")
	router.POST("/createDriver", MiddlewareVehicleDriverDeserialization, vr.handler.CreateVehicleDriver)

}

func MiddlewareVehicleDriverDeserialization(c *gin.Context) {
	var vehicleDriver domain.VehicleDriverCreate

	if err := c.ShouldBindJSON(&vehicleDriver); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to decode JSON"})
		c.Abort()
		return
	}

	c.Set("vehicleDriver", vehicleDriver)
	c.Next()
}
