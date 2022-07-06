package controllers

import (
	"context"
	"log"

	"github.com/Go-MongoDB-REST/database"
	"github.com/Go-MongoDB-REST/helpers"
	"github.com/Go-MongoDB-REST/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repository
type BookController interface {
	FindAll(c *gin.Context)
	Store(c *gin.Context)
}

// Service
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

func (controller BookControllerImp) Store(c *gin.Context) {
	var bookModel models.Book
	if err := helpers.Validate(c, &bookModel); err != 1 {
		return
	}

	book := bson.D{
		{Key: "Name", Value: bookModel.Name},
		{Key: "Type", Value: bookModel.Type},
	}

	res, err := controller.UserCollection.InsertOne(context.TODO(), book)
	if err != nil {
		log.Fatal(err)
	}

	helpers.Response(c, "Insert Book Successfuly", res, 200)
}
