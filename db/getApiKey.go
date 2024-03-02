package db

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetApiKey(discordID string) (string, error) {
	connection := Client.Database("imgHost").Collection("accounts")

	filter := bson.M{"discord_id": discordID}
	var result struct {
		ApiKey string `bson:"api"`
	}
	err := connection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return "", nil
		}
		return "", err
	}
	fmt.Println("d")

	return result.ApiKey, nil
}
