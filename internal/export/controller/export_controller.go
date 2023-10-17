package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tudemaha/ifassion-be/internal/export/services"
	"github.com/tudemaha/ifassion-be/internal/global/responses"
)

func ExportWebHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var response responses.Response

		chainingId := c.Param("id")
		responseData := services.GetResultData(chainingId)

		response.DefaultOK()
		response.Message = "get data success"
		response.Data = responseData
		c.JSON(response.Code, response)
	}
}

func ExportPdfHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// var response responses.Response

		// chainingId := c.Param("id")
		// responseData := services.GetResultData(chainingId)

		// services.CreatePdf(responseData)
	}
}
