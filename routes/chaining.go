package routes

import (
	"github.com/gin-gonic/gin"
	chainingController "github.com/tudemaha/ifassion-be/internal/chaining/controller"
)

func ChainingRoutes(g *gin.RouterGroup) {
	g.POST("", chainingController.NewChainingHandler())
	g.POST("/:id", chainingController.AddIndicatorHandler())
}
