package routes

import (
	"github.com/gin-gonic/gin"
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
	_ = rg.Group("/vehicle")

}
