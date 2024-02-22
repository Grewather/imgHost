package db

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DelFromDb(id string) (string, string, bool) {
	collection := Client.Database("imgHost").Collection("images")
	filter := bson.M{"url": id}

	var image ImageScheme
	err := collection.FindOne(context.TODO(), filter).Decode(&image)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return "", "", false
		}
		panic(err)
	}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	if result.DeletedCount == 0 {
		return "", "", false
	}
	return image.ImgName, image.Extension, true
}
