package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vehicles-service/domain"
	"vehicles-service/handlers"
	"vehicles-service/services"
)

type DriverLicenceRouteHandler struct {
	handler handlers.DriverLicenceHandler
	service services.DriverLicenceService
}

func NewDriverLicenceRouteHandler(handler handlers.DriverLicenceHandler, service services.DriverLicenceService) DriverLicenceRouteHandler {
	return DriverLicenceRouteHandler{handler, service}
}

func (vr *DriverLicenceRouteHandler) DriverLicenceRoute(rg *gin.RouterGroup) {
	router := rg.Group("/driverlicence")
	router.POST("/create", MiddlewareDriverLicenceDeserialization, vr.handler.CreateDriverLicence)
	router.GET("/get/:id", vr.handler.GetLicenceByID)

}

func MiddlewareDriverLicenceDeserialization(c *gin.Context) {
	var driverLicence domain.DriverLicenceCreate

	if err := c.ShouldBindJSON(&driverLicence); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to decode JSON"})
		c.Abort()
		return
	}

	c.Set("driverLicence", driverLicence)
	c.Next()
}
