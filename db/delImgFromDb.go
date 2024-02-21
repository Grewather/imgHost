package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

func DelDbFromDb(accessToken string, id string) bool {
	collection := Client.Database("imgHost").Collection("images")
	filter := bson.M{"url": id}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	if result.DeletedCount == 0 {
		return false
	}
	return true
}
