package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tudemaha/ifassion-be/middleware"
)

func SetupRouterGroup(router *gin.Engine) {
	chaining := router.Group("/chaining", middleware.APIKeyValidator())
	ChainingRoutes(chaining)

	export := router.Group("/export", middleware.APIKeyValidator())
	ExportRoutes(export)
}
