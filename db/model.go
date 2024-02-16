package db

import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	DiscordId string             `bson:"discord_id"`
	Username  string             `bson:"username"`
}

type ImageScheme struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Owner     string             `bson:"owner"`
	ImgName   string             `bson:"img_name"`
	Url       string             `bson:"url"`
	Extension string             `bson:"extension"`
}
