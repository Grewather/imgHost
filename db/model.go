package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	DiscordId string             `bson:"discord_id"`
	Username  string             `bson:"username"`
}
