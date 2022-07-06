package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	Id 		primitive.ObjectID	`json:"id,omitempty"`
	Name 	string 							`json:"name,omitempty" binding:"required"`
	Type 	string 							`json:"type,omitempty" binding:"required"`
}