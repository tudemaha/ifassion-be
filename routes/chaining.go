package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tudemaha/ifassion-be/internal/chaining/controller"
)

func ChainingRoutes(g *gin.RouterGroup) {
	g.POST("", controller.NewChainingHandler())
	g.POST("/:id", controller.AddIndicatorHandler())
}
