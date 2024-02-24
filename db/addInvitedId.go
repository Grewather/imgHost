package db

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddInvitedId(invitedId string) (string, error) {
	collection := Client.Database("imgHost").Collection("invited")

	result := collection.FindOne(context.TODO(), bson.D{{"invitedId", invitedId}})
	if !errors.Is(result.Err(), mongo.ErrNoDocuments) {
		if result.Err() != nil {
			return "", fmt.Errorf("failed to add user to whitelist %v", result.Err())
		}
		return "", fmt.Errorf("User is already whitelisted")
	}

	_, err := collection.InsertOne(context.TODO(), bson.D{{"invitedId", invitedId}})
	if err != nil {
		return "", fmt.Errorf("failed to add user to whitelist: %v", err)
	}
	return "User successfully added to whitelist", nil
}
