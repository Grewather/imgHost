package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetImagesToLoad(id string, pageSize, pageNum int) ([]ImageScheme, error) {
	connection := Client.Database("imgHost").Collection("images")
	filter := bson.M{"owner": id}
	opt := options.Find().SetLimit(int64(pageSize)).SetSkip(int64((pageNum - 1) * pageSize))
	cur, err := connection.Find(context.TODO(), filter, opt)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var results []ImageScheme

	for cur.Next(context.Background()) {
		var result ImageScheme
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("No results found")
	}

	return results, nil
}
