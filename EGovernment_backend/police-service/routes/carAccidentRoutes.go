package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"police-service/domain"
	"police-service/handlers"
	"police-service/services"
)

type CarAccidentRouteHandler struct {
	handler handlers.CarAccidentHandler
	service services.CarAccidentService
}

func NewCarAccidentRouteHandler(handler handlers.CarAccidentHandler, service services.CarAccidentService) CarAccidentRouteHandler {
	return CarAccidentRouteHandler{handler, service}
}

func (d *CarAccidentRouteHandler) CarAccidentRoute(rg *gin.RouterGroup) {
	router := rg.Group("/carAccident")
	router.POST("/createCarAccident", MiddlewareCarAccidentDeserialization, d.handler.CreateCarAccident)
	router.GET("/all", d.handler.GetAllCarAccidents)
	router.GET("/get/:id", d.handler.GetCarAccidentByID)
	router.GET("/getPolicemanCarAccident", d.handler.GetCarAccidentsByPolicemanID)
	router.GET("/getDriverCarAccident", d.handler.GetCarAccidentsByDriver)
	router.GET("/get/carAccidentType/:carAccidentType", d.handler.GetAllCarAccidentsByType)
	router.GET("/get/degreeOfAccident/:degreeOfAccident", d.handler.GetAllCarAccidentsByDegree)
	router.GET("/get/carAccidentType/:carAccidentType/year/:year", d.handler.GetAllCarAccidentsByTypeAndYear)
	router.GET("/get/degreeOfAccident/:degreeOfAccident/year/:year", d.handler.GetAllCarAccidentsByDegreeAndYear)

}

func MiddlewareCarAccidentDeserialization(c *gin.Context) {
	var carAccident domain.CarAccidentCreate

	if err := c.ShouldBindJSON(&carAccident); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to decode JSON"})
		c.Abort()
		return
	}

	c.Set("carAccident", carAccident)
	c.Next()
}
