package controller

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tudemaha/ifassion-be/internal/chaining/dto"
	"github.com/tudemaha/ifassion-be/internal/global/responses"
	"github.com/tudemaha/ifassion-be/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
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

		client = mongo.MongoConnection("indicators")
		var indicator dto.Indicator
		filter := bson.D{{Key: "code", Value: "I01"}}
		if err := client.Coll.FindOne(context.TODO(), filter).Decode(&indicator); err != nil {
			response.DefaultInternalError()
			c.AbortWithStatusJSON(response.Code, response)
			return
		}

		response.DefaultCreated()
		response.Message = "new chaining inserted successfully"
		chainingData := map[string]interface{}{
			"id":      result.InsertedID,
			"fninish": false,
		}
		questionData := map[string]string{
			"id":       indicator.Code,
			"question": "Apakah Anda " + indicator.IndicatorString + "?",
		}
		response.Data = map[string]interface{}{
			"chaining": chainingData,
			"question": questionData,
		}
		c.JSON(response.Code, response)
	}
}
