package controller

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/tudemaha/ifassion-be/internal/export/utils"
	"github.com/tudemaha/ifassion-be/internal/global/dto"
	"github.com/tudemaha/ifassion-be/internal/global/responses"
	"github.com/tudemaha/ifassion-be/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ExportWebHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var response responses.Response

		chainingId := c.Param("id")

		clientResult := mongo.MongoConnection("results")
		defer mongo.MongoCloseConnection(clientResult)

		id, _ := primitive.ObjectIDFromHex(chainingId)
		filter := bson.D{{Key: "_id", Value: id}}

		var result dto.ResultData
		if err := clientResult.Coll.FindOne(context.TODO(), filter).Decode(&result); err != nil {
			panic(err)
		}

		result.Database.True = utils.SliceWithUniqueValues(result.Database.True)
		clientIndicator := mongo.MongoConnection("indicators")
		defer mongo.MongoCloseConnection(clientIndicator)

		filter = bson.D{{
			Key: "code",
			Value: bson.D{{
				Key:   "$in",
				Value: result.Database.True,
			}},
		}}
		cursor, err := clientIndicator.Coll.Find(context.TODO(), filter)
		if err != nil {
			panic(err)
		}

		var indicators []string
		for cursor.Next(context.TODO()) {
			var indicator dto.Indicator
			if err := cursor.Decode(&indicator); err != nil {
				panic(err)
			}
			indicators = append(indicators, indicator.IndicatorString)
		}

		clientPassion := mongo.MongoConnection("passions")
		defer mongo.MongoCloseConnection(clientPassion)

		var passion dto.Passion
		filter = bson.D{{Key: "code", Value: result.Passion}}
		if err := clientPassion.Coll.FindOne(context.TODO(), filter).Decode(&passion); err != nil {
			panic(err)
		}

		response.DefaultOK()
		response.Message = "get data success"
		response.Data = map[string]interface{}{
			"id":         chainingId,
			"indicators": indicators,
			"passion":    passion.Interest,
			"time":       result.Time[:19],
		}
		c.JSON(response.Code, response)
	}
}
