package database

import (
	"context"
	"errors"
	"fmt"
	"notes-back/helpers"
	"notes-back/types/requestTypes"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetItems(collection *mongo.Collection, query requestTypes.GetQuery) (*mongo.Cursor, int, error) {
	const defaultLimit = 100

	err := validateQuery(&query)
	if err != nil {
		return nil, 400, err
	}

	findOptions := options.Find().SetLimit(defaultLimit)
	var queryFilter bson.D
	queryProjection := bson.D{}

	if len(query.Return) > 0 {
		for _, r := range query.Return {
			queryProjection = append(queryProjection, bson.E{Key: r, Value: 1})
		}
		findOptions.SetProjection(queryProjection)
	}

	if len(query.Filters) == 1 {
		queryFilter = bson.D{{Key: query.Filters[0].Key, Value: bson.D{{Key: query.Filters[0].Comparison, Value: query.Filters[0].Value}}}}
	} else {
		filters := bson.A{}
		for _, f := range query.Filters {
			filters = append(filters, bson.D{{Key: f.Key, Value: bson.D{{Key: f.Comparison, Value: f.Value}}}})
		}
		queryFilter = bson.D{{Key: query.LogicalOperator, Value: filters}}
	}

	cursor, err := collection.Find(context.Background(), queryFilter, findOptions)
	if err != nil {
		return nil, 400, err
	}
	return cursor, 200, nil
}

func validateQuery(query *requestTypes.GetQuery) error {
	if len(query.Filters) > 1 && query.LogicalOperator == "" {
		return errors.New("missing logical operator")
	}
	for _, f := range query.Filters {
		if !helpers.Contains(Comparisons, f.Comparison) {
			return fmt.Errorf("invalid comparison operator '%s'", f.Comparison)
		}
	}
	return nil
}


