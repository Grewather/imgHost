package db

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetImg(id string) (string, string, error) {
	collection := Client.Database("imgHost").Collection("images")
	filter := bson.M{"url": id}
	var imageUrl ImageScheme
	err := collection.FindOne(context.TODO(), filter).Decode(&imageUrl)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return "", "", errors.New("image not found")
		}
		return "", "", err
	}
	fmt.Println(imageUrl)
	ext := imageUrl.ImgName + imageUrl.Extension
	return imageUrl.Owner, ext, err
}
