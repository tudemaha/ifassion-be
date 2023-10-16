package controller

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tudemaha/ifassion-be/internal/chaining/dto"
	"github.com/tudemaha/ifassion-be/internal/global/responses"
	"github.com/tudemaha/ifassion-be/pkg/mongo"
)

func NewChainingHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var response responses.Response

		var newData dto.ResultData
		newData.Time = time.Now().String()
		newData.Status = false

		client := mongo.MongoConnection("results")
		defer mongo.MongoCloseConnection(client)

		result, err := client.Coll.InsertOne(context.TODO(), newData)
		if err != nil {
			response.DefaultBadRequest()
			c.AbortWithStatusJSON(response.Code, response)
			return
		}

		response.DefaultCreated()
		response.Message = "new chaining inserted successfully"
		response.Data = map[string]interface{}{"id": result.InsertedID}
		c.JSON(response.Code, response)
	}
}
