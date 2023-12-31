package controller

import (
	"context"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	chainingDto "github.com/tudemaha/ifassion-be/internal/chaining/dto"
	"github.com/tudemaha/ifassion-be/internal/chaining/utils"
	globalDto "github.com/tudemaha/ifassion-be/internal/global/dto"
	"github.com/tudemaha/ifassion-be/internal/global/responses"
	"github.com/tudemaha/ifassion-be/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewChainingHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var response responses.Response

		var newData globalDto.ResultData
		local, _ := time.LoadLocation("Asia/Makassar")
		newData.Time = time.Now().In(local).Format("2006-01-02 15:04:05") + " WITA"
		newData.Database.False = make([]string, 0)
		newData.Database.True = make([]string, 0)
		newData.Rules = make([]string, 0)
		newData.Status = false

		clientResult := mongo.MongoConnection("results")
		defer mongo.MongoCloseConnection(clientResult)

		result, err := clientResult.Coll.InsertOne(context.TODO(), newData)
		if err != nil {
			response.DefaultInternalError()
			response.Data = map[string]string{"error": err.Error()}
			c.AbortWithStatusJSON(response.Code, response)
			return
		}

		clientIndicator := mongo.MongoConnection("indicators")
		defer mongo.MongoCloseConnection(clientIndicator)

		var indicator globalDto.Indicator
		filter := bson.D{{Key: "code", Value: "I01"}}
		if err := clientIndicator.Coll.FindOne(context.TODO(), filter).Decode(&indicator); err != nil {
			response.DefaultInternalError()
			response.Data = map[string]string{"error": err.Error()}
			c.AbortWithStatusJSON(response.Code, response)
			return
		}

		response.DefaultCreated()
		response.Message = "new chaining inserted successfully"
		chainingData := map[string]interface{}{
			"id":     result.InsertedID,
			"finish": false,
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

func AddIndicatorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var response responses.Response

		chainingId := c.Param("id")
		var QuestionInput chainingDto.UserIndicatorInput
		if err := c.BindJSON(&QuestionInput); err != nil {
			response.DefaultInternalError()
			response.Data = map[string]string{"error": err.Error()}
			c.AbortWithStatusJSON(response.Code, response)
			return
		}
		statusString := strconv.FormatBool(QuestionInput.QuestionStatus)

		clientResult := mongo.MongoConnection("results")
		defer mongo.MongoCloseConnection(clientResult)

		id, _ := primitive.ObjectIDFromHex(chainingId)
		filterID := bson.D{{Key: "_id", Value: id}}
		update := bson.D{{
			Key: "$addToSet",
			Value: bson.D{{
				Key:   "database." + statusString,
				Value: QuestionInput.QuestionID,
			}},
		}}
		_, err := clientResult.Coll.UpdateOne(context.TODO(), filterID, update)
		if err != nil {
			response.DefaultNotFound()
			response.Data = map[string]string{"error": err.Error()}
			c.AbortWithStatusJSON(response.Code, response)
			return
		}

		var resultData globalDto.ResultData
		if err := clientResult.Coll.FindOne(context.TODO(), filterID).Decode(&resultData); err != nil {
			response.DefaultNotFound()
			response.Data = map[string]string{"error": err.Error()}
			c.AbortWithStatusJSON(response.Code, response)
			return
		}

		clientRules := mongo.MongoConnection("rules")
		defer mongo.MongoCloseConnection(clientRules)

		cursor, err := clientRules.Coll.Find(context.TODO(), bson.D{})
		if err != nil {
			response.DefaultInternalError()
			response.Data = map[string]string{"error": err.Error()}
			c.AbortWithStatusJSON(response.Code, response)
			return
		}
		defer cursor.Close(context.TODO())

		var ruleData globalDto.Rule
		for cursor.Next(context.TODO()) {
			if err := cursor.Decode(&ruleData); err != nil {
				response.DefaultInternalError()
				response.Data = map[string]string{"error": err.Error()}
				c.AbortWithStatusJSON(response.Code, response)
				return
			}

			if utils.AllElementExist(resultData.Database.True, ruleData.If) {
				if utils.ContainsElement(resultData.Rules, ruleData.Code) {
					continue
				}
				update = bson.D{
					{
						Key: "$addToSet",
						Value: bson.D{{
							Key:   "rules",
							Value: ruleData.Code,
						}},
					}}
				_, err = clientResult.Coll.UpdateOne(context.TODO(), filterID, update)
				if err != nil {
					response.DefaultNotFound()
					response.Data = map[string]string{"error": err.Error()}
					c.AbortWithStatusJSON(response.Code, response)
					return
				}

				clientPassion := mongo.MongoConnection("passions")
				defer mongo.MongoCloseConnection(clientPassion)
				if ruleData.Then[0] == 'P' {
					filter := bson.D{{Key: "code", Value: ruleData.Then}}
					var passion globalDto.Passion
					if err := clientPassion.Coll.FindOne(context.TODO(), filter).Decode(&passion); err != nil {
						response.DefaultInternalError()
						response.Data = map[string]string{"error": err.Error()}
						c.AbortWithStatusJSON(response.Code, response)
						return
					}

					update = bson.D{{
						Key: "$set",
						Value: bson.D{{
							Key:   "status",
							Value: true,
						}, {
							Key:   "passion",
							Value: ruleData.Then,
						}},
					}}
					_, err = clientResult.Coll.UpdateOne(context.TODO(), filterID, update)
					if err != nil {
						response.DefaultInternalError()
						response.Data = map[string]string{"error": err.Error()}
						c.AbortWithStatusJSON(response.Code, response)
						return
					}

					response.DefaultOK()
					response.Message = "chaining finish"
					chainingData := map[string]interface{}{
						"id":     id,
						"finish": true,
					}
					response.Data = map[string]interface{}{
						"chaining": chainingData,
					}
					c.JSON(response.Code, response)
					return
				}

				if ruleData.Then[0] == 'I' {
					update = bson.D{{
						Key: "$addToSet",
						Value: bson.D{{
							Key:   "database.true",
							Value: ruleData.Then,
						}},
					}}
					_, err := clientResult.Coll.UpdateOne(context.TODO(), filterID, update)
					if err != nil {
						response.DefaultInternalError()
						response.Data = map[string]string{"error": err.Error()}
						c.AbortWithStatusJSON(response.Code, response)
						return
					}
				}
				break
			}
		}

		// next question
		filterIf := bson.D{{
			Key: "if",
			Value: bson.D{{
				Key:   "$in",
				Value: resultData.Database.True,
			}},
		}}

		cursor, err = clientRules.Coll.Find(context.TODO(), filterIf)
		if err != nil {
			response.DefaultInternalError()
			response.Data = map[string]string{"error": err.Error()}
			c.AbortWithStatusJSON(response.Code, response)
			return
		}
		defer cursor.Close(context.TODO())

		countDatabase := len(resultData.Database.False) + len(resultData.Database.True)
		if countDatabase == 22 {
			update = bson.D{{
				Key: "$set",
				Value: bson.D{{
					Key:   "status",
					Value: true,
				}, {
					Key:   "passion",
					Value: "P99",
				}},
			}}
			_, err = clientResult.Coll.UpdateOne(context.TODO(), filterID, update)
			if err != nil {
				response.DefaultInternalError()
				response.Data = map[string]string{"error": err.Error()}
				c.AbortWithStatusJSON(response.Code, response)
				return
			}

			response.DefaultOK()
			response.Message = "chaining finish"
			chainingData := map[string]interface{}{
				"id":     id,
				"finish": true,
			}
			response.Data = map[string]interface{}{
				"chaining": chainingData,
			}
			c.JSON(response.Code, response)
			return
		}

		var nextIndicatorID string
		nextIndicatorID = utils.IncrementCode(resultData.Database.True, resultData.Database.False)
		for cursor.Next(context.TODO()) {
			if err := cursor.Decode(&ruleData); err != nil {
				response.DefaultInternalError()
				response.Data = map[string]string{"error": err.Error()}
				c.AbortWithStatusJSON(response.Code, response)
				return
			}

			if utils.ContainsElement(resultData.Rules, ruleData.Code) {
				continue
			}

			indicator, status := utils.OneUniqueElement(resultData.Database.True, resultData.Database.False, ruleData.If)
			if status {
				nextIndicatorID = indicator
				break
			}
		}

		clientIndicator := mongo.MongoConnection("indicators")
		defer mongo.MongoCloseConnection(clientIndicator)

		filterIndicator := bson.D{{
			Key:   "code",
			Value: nextIndicatorID,
		}}
		var indicator globalDto.Indicator
		if err := clientIndicator.Coll.FindOne(context.TODO(), filterIndicator).Decode(&indicator); err != nil {
			response.DefaultInternalError()
			response.Data = map[string]string{"error": err.Error()}
			c.AbortWithStatusJSON(response.Code, response)
			return
		}

		response.DefaultOK()
		response.Message = "chaining question inserted successfully"
		chainingData := map[string]interface{}{
			"id":     id,
			"finish": false,
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
