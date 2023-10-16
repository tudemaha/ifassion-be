package mongo

import "go.mongodb.org/mongo-driver/mongo"

type Mongo struct {
	Client *mongo.Client
	Coll   *mongo.Collection
}
