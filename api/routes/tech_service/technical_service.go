package tech_service

import (
	"github.com/anthophora/tamircity/api/handler"
	"github.com/gin-gonic/gin"
)

func TechnicalServiceRouter(router *gin.Engine, technicalServiceHandler handler.TechnicalServiceHandler) {
	technicalServiceRoute := router.Group("api/v1/tech/technical-services")
	{
		technicalServiceRoute.GET("", technicalServiceHandler.GetAll)
		technicalServiceRoute.GET("/query", technicalServiceHandler.GetAllByFilter)
		technicalServiceRoute.POST("", technicalServiceHandler.Create)
		technicalServiceRoute.GET("/:id", technicalServiceHandler.Get)
		technicalServiceRoute.PUT("/:id", technicalServiceHandler.Update)
		technicalServiceRoute.DELETE("/:id", technicalServiceHandler.Delete)
	}
}