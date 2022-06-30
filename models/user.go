package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID 			primitive.ObjectID 	`json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName 	string 				`json:"firstname,omitempty" bson:"firstname"`
	LastName 	string 				`json:"lastname,omitempty" bson:"lastname"`
	Email 		string 				`json:"email,omitempty" bson:"email"`
	Password 	string 				`json:"password,omitempty" bson:"password"`
}