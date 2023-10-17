package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type DatabaseInsert struct {
	True  []string `bson:"true"`
	False []string `bson:"false"`
}

type ResultData struct {
	Time     string         `bson:"time"`
	Database DatabaseInsert `bson:"database"`
	Rules    []string       `bson:"rules"`
	Status   bool           `bson:"status"`
	Passion  string         `bson:"passion"`
}

type Indicator struct {
	ID              primitive.ObjectID `bson:"_id"`
	Code            string             `bson:"code"`
	IndicatorString string             `bson:"indicator"`
}

type Rule struct {
	ID   primitive.ObjectID `bson:"_id"`
	Code string             `bson:"code"`
	If   []string           `bson:"if"`
	Then string             `bson:"then"`
}

type Passion struct {
	ID       primitive.ObjectID `bson:"_id"`
	Code     string             `bson:"code"`
	Interest string             `bson:"interest"`
}
