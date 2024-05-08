package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vehicles-service/domain"
	"vehicles-service/handlers"
	"vehicles-service/services"
)

type VehicleRouteHandler struct {
	handler handlers.VehicleHandler
	service services.VehicleService
}

func NewVehicleRouteHandler(handler handlers.VehicleHandler, service services.VehicleService) VehicleRouteHandler {
	return VehicleRouteHandler{handler, service}
}

func (vr *VehicleRouteHandler) VehicleRoute(rg *gin.RouterGroup) {
	router := rg.Group("/vehicle")
	router.POST("/createDriver", MiddlewareVehicleDeserialization, vr.handler.CreateVehicleDriver)

}

func MiddlewareVehicleDeserialization(c *gin.Context) {
	var vehicleDriver domain.VehicleDriverCreate

	if err := c.ShouldBindJSON(&vehicleDriver); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to decode JSON"})
		c.Abort()
		return
	}

	c.Set("vehicleDriver", vehicleDriver)
	c.Next()
}
