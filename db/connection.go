package db

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var Client *mongo.Client

func ConnectToDb() {
	godotenv.Load(".env")
	serverApi := options.ServerAPI(options.ServerAPIVersion1)
	// use here env variable
	opts := options.Client().ApplyURI(os.Getenv("MONGO_URL")).SetServerAPIOptions(serverApi)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {

		log.Fatal(err)
	}
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {

		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	Client = client
}
