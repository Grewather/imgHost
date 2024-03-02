package db

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"imgHost/models"
)

func GetApiKey(api string) (models.Account, error) {
	collection := Client.Database("imgHost").Collection("accounts")
	filter := bson.M{"api": api}
	var result models.Account
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			fmt.Println(err)
			return models.Account{}, err
		}
		fmt.Println(err)
		return models.Account{}, err
	}
	fmt.Println(result)

	return result, nil
}
