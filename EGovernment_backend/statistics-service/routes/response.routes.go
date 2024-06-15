package routes

import (
	"statistics-service/handlers"
	"statistics-service/services"

	"github.com/gin-gonic/gin"
)

type ResponseRouteHandler struct {
	handler handlers.ResponseHandler
	service services.ResponseService
}

func NewResponseRouteHandler(handler handlers.ResponseHandler, service services.ResponseService) ResponseRouteHandler {
	return ResponseRouteHandler{handler, service}
}

func (rc *ResponseRouteHandler) Route(rg *gin.RouterGroup) {
	router := rg.Group("/response")
	router.POST("/create", rc.handler.Create)
	router.GET("/all", rc.handler.GetAll)
	router.GET("/get/:id", rc.handler.GetByID)
}
