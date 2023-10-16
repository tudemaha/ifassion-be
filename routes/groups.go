package routes

import "github.com/gin-gonic/gin"

func SetupRouterGroup(router *gin.Engine) {
	chaining := router.Group("/chaining")
	ChainingRoutes(chaining)
}
