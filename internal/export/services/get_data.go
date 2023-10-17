package services

import (
	"context"

	exportDto "github.com/tudemaha/ifassion-be/internal/export/dto"
	"github.com/tudemaha/ifassion-be/internal/export/utils"
	globalDto "github.com/tudemaha/ifassion-be/internal/global/dto"
	"github.com/tudemaha/ifassion-be/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetResultData(chainingId string) exportDto.ResponseData {
	clientResult := mongo.MongoConnection("results")
	defer mongo.MongoCloseConnection(clientResult)

	id, _ := primitive.ObjectIDFromHex(chainingId)
	filter := bson.D{{Key: "_id", Value: id}}

	var result globalDto.ResultData
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
	defer cursor.Close(context.TODO())

	var indicators []string
	for cursor.Next(context.TODO()) {
		var indicator globalDto.Indicator
		if err := cursor.Decode(&indicator); err != nil {
			panic(err)
		}
		indicators = append(indicators, indicator.IndicatorString)
	}

	clientPassion := mongo.MongoConnection("passions")
	defer mongo.MongoCloseConnection(clientPassion)

	var passion globalDto.Passion
	filter = bson.D{{Key: "code", Value: result.Passion}}
	if err := clientPassion.Coll.FindOne(context.TODO(), filter).Decode(&passion); err != nil {
		panic(err)
	}

	return exportDto.ResponseData{
		ID:         chainingId,
		Indicators: indicators,
		Passion:    passion.Interest,
		Time:       result.Time[:19],
	}
}
