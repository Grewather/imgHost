package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct {
	ID        string `bson:"_id,omitempty" json:"_id"`
	DiscordId string `bson:"discord_id" json:"id"`
	Username  string `bson:"username"`
	ApiKey    string `bson:"api"`
}

type ImageScheme struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Owner     string             `bson:"owner"`
	ImgName   string             `bson:"img_name"`
	Url       string             `bson:"url"`
	Extension string             `bson:"extension"`
}
