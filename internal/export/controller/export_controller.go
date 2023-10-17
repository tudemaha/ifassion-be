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
		var response responses.Response

		chainingId := c.Param("id")
		responseData := services.GetResultData(chainingId)

		filename, err := services.CreatePdf(responseData)
		if err != nil {
			response.DefaultInternalError()
			response.Data = map[string]string{"error": err.Error()}
			c.AbortWithStatusJSON(response.Code, response)
			return
		}

		protocol := "http://"
		if c.Request.TLS != nil {
			protocol = "https://"
		}
		file := protocol + c.Request.Host + "/pdf/" + filename
		response.DefaultOK()
		response.Message = "pdf generated successfully"
		response.Data = map[string]string{
			"pdf": file,
		}
		c.JSON(response.Code, response)
	}
}
