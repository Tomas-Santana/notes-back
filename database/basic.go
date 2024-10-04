package database

import (
	"context"
	"errors"

	"notes-back/helpers"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCollection(db *mongo.Database, collectionName string) *mongo.Collection {
	return db.Collection(collectionName)
}

func GetItemByID(collection *mongo.Collection, id string) (*mongo.SingleResult, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err

	}
	filter := bson.D{{Key: "_id", Value: objectID}}
	return collection.FindOne(context.Background(), filter), nil
}


func UpdateWithID(collection *mongo.Collection, id string, updateFields bson.D, operator string) (*mongo.UpdateResult, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	if !helpers.Contains(UpdateOperators, operator) {
		return nil, errors.New("invalid operator")
	}

	update := bson.D{{Key: operator, Value: updateFields}}
	result, err := collection.UpdateByID(context.Background(), objectID, update)

	if err != nil {
		return nil, err
	}

	return result, nil

}


