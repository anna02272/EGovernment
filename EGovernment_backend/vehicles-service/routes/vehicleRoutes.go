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
	router.POST("/createVehicle", MiddlewareVehicleDeserialization, vr.handler.CreateVehicle)

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
