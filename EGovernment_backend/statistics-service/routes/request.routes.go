package routes

import (
	"statistics-service/handlers"
	"statistics-service/services"

	"github.com/gin-gonic/gin"
)

type RequestRouteHandler struct {
	handler handlers.RequestHandler
	service services.RequestService
}

func NewRequestRouteHandler(handler handlers.RequestHandler, service services.RequestService) RequestRouteHandler {
	return RequestRouteHandler{handler, service}
}

func (rc *RequestRouteHandler) Route(rg *gin.RouterGroup) {
	router := rg.Group("/request")
	router.POST("/create", rc.handler.Create)
	router.GET("/all", rc.handler.GetAll)
	router.GET("/get/:id", rc.handler.GetByID)
}
