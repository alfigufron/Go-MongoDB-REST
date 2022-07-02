package utils

import (
	"context"
	"os"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InitMongoDB() {
	uri := os.Getenv("MONGODB_URI")

	log.Info("Connecting to MongoDB Database!")
	dbClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	log.Info("Pinging to MongoDB Database!")
	if err := dbClient.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	defer func() {
		if err := dbClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	log.Info("Connected to Database Successfully!")
}
