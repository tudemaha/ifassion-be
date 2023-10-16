package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type DatabaseInsert struct {
	True  []string `bson:"true"`
	False []string `bson:"false"`
}

type ResultData struct {
	Time     string         `bson:"time"`
	Database DatabaseInsert `bson:"database"`
	Status   bool           `bson:"status"`
}

type Indicator struct {
	ID              primitive.ObjectID `bson:"_id"`
	Code            string             `bson:"code"`
	IndicatorString string             `bson:"indicator"`
}
