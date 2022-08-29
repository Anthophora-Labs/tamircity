package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafakocatepe/Tamircity/api/handler"
)

func BrandRouter(router *gin.Engine, brandHandler handler.BrandHandler) {
	route := router.Group("api/v1/brands")
	{
		route.GET("/", brandHandler.GetAll)
		route.GET("/query", brandHandler.GetAllByDeviceTypeId)
		route.POST("/", brandHandler.Create)
	}
}
