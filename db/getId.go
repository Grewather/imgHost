package db

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"os"
)

func GetIdFromDb(Id string) bool {
	godotenv.Load(".env")
	collection := Client.Database("imgHost").Collection("invited")
	filter := bson.M{"invitedId": Id}
	if Id == os.Getenv("ADMIN_ID") {
		return true
	}
	var result bson.M
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return false
	}
	fmt.Print(result)
	if result != nil {
		return true
	} else {
		return false
	}
}
