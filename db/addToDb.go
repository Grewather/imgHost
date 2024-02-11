package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddToDb(DiscordId, Username string) {
	collection := Client.Database("imgHost").Collection("accounts")

	filter := bson.D{{"discord_id", DiscordId}}
	update := bson.D{{"$set", bson.D{
		{"username", Username},
	}}}

	opts := options.Update().SetUpsert(true)
	_, err := collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		panic(err)
	}
}
