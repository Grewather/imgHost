package db

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"imgHost/utils"
)

func AddImgToDb(imgName, owner string) {

	randString := utils.GetRandomString()
	for {
		if checkIfYouCanAdd(randString) {
			break
		}
		randString = utils.GetRandomString()
	}

	collection := Client.Database("imgHost").Collection("images")
	_, err := collection.InsertOne(context.TODO(), bson.D{
		{"owner", owner},
		{"img_name", imgName},
		{"url", randString},
	})
	if err != nil {
		panic(err)
	}

	return

}

func checkIfYouCanAdd(randString string) bool {
	collection := Client.Database("imgHost").Collection("images")
	filter := bson.M{"url": randString}
	var imageUrl ImageScheme
	err := collection.FindOne(context.TODO(), filter).Decode(&imageUrl)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return true
		}
		panic(err)
	}
	return false
}
