package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
	ID 			primitive.ObjectID 	`json:"_id,omitempty" bson:"_id,omitempty"`
	Name		string				`json:"name,omitempty" bson:"name"`
	Director	string				`json:"director,omitempty" bson:"director"`
	Watched		bool				`json:"watched,omitempty" bson:"watched"`
}