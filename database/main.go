package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoClient(uri string) *mongo.Client {

	client, err := mongo.Connect(context.TODO(), options.Client().
	ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	return client
}
