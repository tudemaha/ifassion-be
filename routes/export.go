package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tudemaha/ifassion-be/internal/export/controller"
)

func ExportRoutes(g *gin.RouterGroup) {
	g.GET("/web/:id", controller.ExportWebHandler())
}
