package routes

import (
	"court-service/handlers"
	"court-service/services"
	"github.com/gin-gonic/gin"
)

type SubjectRouteHandler struct {
	handler handlers.SubjectHandler
	service services.SubjectService
}

func NewSubjectRouteHandler(handler handlers.SubjectHandler, service services.SubjectService) SubjectRouteHandler {
	return SubjectRouteHandler{handler, service}
}

func (sr *SubjectRouteHandler) SubjectRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/subject")
	router.POST("/create", sr.handler.CreateSubject)
}
