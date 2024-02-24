package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func RemoveId(id string) bool {
	collection := Client.Database("imgHost").Collection("invited")
	filter := bson.M{"invitedId": id}

	result := collection.FindOne(context.TODO(), filter)
	if result.Err() == mongo.ErrNoDocuments {
		return false
	} else if result.Err() != nil {
		return false
	}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return false
	}

	if deleteResult.DeletedCount == 0 {
		return false
	}

	return true
}
