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
	router.GET("/get/:id", sr.handler.GetDelict)
	router.GET("/subjects/:id", sr.handler.GetSubject)
	router.GET("/subjects", sr.handler.GetAllSubjects)
	router.PUT("/subjects/:id/status", sr.handler.UpdateSubjectStatus)
	router.PUT("/subjects/:id/judgment", sr.handler.UpdateSubjectJudgment)
	router.PUT("/subjects/:id/compromis", sr.handler.UpdateSubjectCompromis)

}
