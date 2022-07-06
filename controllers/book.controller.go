package controllers

import (
	"context"
	"log"

	"github.com/Go-MongoDB-REST/database"
	"github.com/Go-MongoDB-REST/helpers"
	"github.com/Go-MongoDB-REST/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repository
type BookController interface {
	FindAll(c *gin.Context)
	Detail(c *gin.Context)
	Store(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
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

	var results []bson.M
	
	for cursor.Next(context.TODO()) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}

	helpers.Response(c, "Get List Book Successfully", results, 200)
}

func (controller BookControllerImp) Detail(c *gin.Context)  {
	// FindOne
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		helpers.Response(c, "Invalid ID", nil, 400)
		return
	}

	filter := bson.D{{Key: "_id", Value: id}}

	var result bson.M
	err = controller.UserCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
		helpers.Response(c, "Get Book Failed", nil, 500)
		return
	}

	helpers.Response(c, "Get Book Successfuly", result, 200)
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

func (controller BookControllerImp) Update(c *gin.Context)  {
	var bookModel models.Book
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		helpers.Response(c, "Invalid ID", nil, 400)
		return
	}

	if err := helpers.Validate(c, &bookModel); err != 1 {
		return
	}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "Name", Value: bookModel.Name}, 
			{Key: "Type", Value: bookModel.Type},
		},
	}}

	res, err := controller.UserCollection.UpdateByID(context.TODO(), id, update)
	
	if err != nil {
		log.Fatal(err)
		helpers.Response(c, "Update Book Failed", nil, 500)
		return
	}

	helpers.Response(c, "Update Book Successfuly", res, 200)
}

func (controller BookControllerImp) Delete(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		helpers.Response(c, "Invalid ID", nil, 400)
		return
	}

	filter := bson.D{{Key: "_id", Value: id}}

	res, err := controller.UserCollection.DeleteOne(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
		helpers.Response(c, "Delete Book Failed", nil, 500)
		return
	}

	helpers.Response(c, "Delete Book Successfuly", res, 200)
}