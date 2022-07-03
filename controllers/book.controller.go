package controllers

import (
	"context"
	"log"

	"github.com/Go-MongoDB-REST/database"
	"github.com/Go-MongoDB-REST/helpers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookController interface {
	FindAll(c *gin.Context)
}

type BookControllerImp struct {
	UserCollection *mongo.Collection
}

func NewBookController() BookController {
	client := database.DB
	userCollection := database.GetCollection(client, "books")

	return &BookControllerImp{UserCollection: userCollection}
}

func (controller BookControllerImp) FindAll(c *gin.Context) {
	cursor, err := controller.UserCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(c)

	helpers.Response(c, "Get List Book Successfully", nil, 200)
}
