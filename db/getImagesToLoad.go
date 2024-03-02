package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"imgHost/models"
)

func GetImagesToLoad(id string) ([]models.ImageScheme, error) {
	connection := Client.Database("imgHost").Collection("images")
	filter := bson.M{"owner": id}
	cur, err := connection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var results []models.ImageScheme

	for cur.Next(context.Background()) {
		var result models.ImageScheme
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("no results found")
	}

	return results, nil
}
