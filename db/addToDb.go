package db

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"imgHost/models"
	"imgHost/utils"
)

func AddToDb(account models.Account) {
	collection := Client.Database("imgHost").Collection("accounts")

	filter := bson.M{"discord_id": account.DiscordId}

	var result models.Account
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			for {
				account.ApiKey = utils.GetRandomString()
				apiKeyFilter := bson.M{"api": account.ApiKey}
				err := collection.FindOne(context.TODO(), apiKeyFilter).Decode(&result)
				if errors.Is(err, mongo.ErrNoDocuments) {
					break
				}
			}
			_, err := collection.InsertOne(context.TODO(), account)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	} else {
		update := bson.M{"$set": bson.M{"username": account.Username}}
		_, err := collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			panic(err)
		}
	}
}
