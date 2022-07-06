package database

import (
	"context"
	"os"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DB * mongo.Client

func InitMongoDB() {
	MONGODB_URI := os.Getenv("MONGODB_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI(MONGODB_URI))
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Connecting to MongoDB Client!")
	if err := client.Connect(context.TODO()); err != nil {
		log.Fatal(err)
	}

	log.Info("Pinging to MongoDB Client!")
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	log.Info("Connected to Database Successfully!")
	DB = client
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	MONGODB_DB := os.Getenv("MONGODB_DB")
	if MONGODB_DB == "" {
		panic("MONGODB_DB is not set")
	}
	collection := client.Database(MONGODB_DB).Collection(collectionName)

	return collection
}