package routes

import (
	"court-service/domain"
	"court-service/handlers"
	"court-service/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CitizenRouteHandler struct {
	handler handlers.CitizenHandler
	service services.CitizenService
}

// Ispravljamo da vrati pokazivač na CitizenRouteHandler
func NewCitizenRouteHandler(handler handlers.CitizenHandler, service services.CitizenService) CitizenRouteHandler {
	return CitizenRouteHandler{handler, service}
}

func (cr *CitizenRouteHandler) CitizenRoute(rg *gin.RouterGroup) {
	router := rg.Group("/citizen")
	router.Use(handlers.ExtractTraceInfoMiddleware())

	router.POST("/create", cr.handler.AddCitizen)
	router.GET("/all", cr.handler.GetAllCitizens)
	router.GET("/get/:jmbg", cr.handler.GetCitizenByID)

}

func MiddlewareCitizenDeserialization(c *gin.Context) {
	var citizen domain.Citizen

	if err := c.ShouldBindJSON(&citizen); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to decode JSON"})
		c.Abort()
		return
	}

	c.Set("citizen", citizen)
	c.Next()
}
