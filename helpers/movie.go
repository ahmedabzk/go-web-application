package helpers

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"

	// "encoding/json"
	"log"

	"github.com/ahmed/go-web/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertOneMovie insert one movie to the database
func InsertOneMovie(movie models.Movie) interface{} {
	inserted, err := Collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal("could not insert a movie to the database")
	}
	return inserted.InsertedID
}

// SearchForAMovie search for a movie by id
func SearchForAMovie(movieId string) interface{} {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	result := Collection.FindOne(context.Background(), filter)

	return result
}

// DeleteOneMovie delete one movie
func DeleteOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	result, err := Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.DeletedCount)
}

// DeleteAllMovies delete all movies
func DeleteAllMovies() {
	result, err := Collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.DeletedCount)
	return
}

// DisplayAllMovies display all movies
func DisplayAllMovies() []primitive.M {

	filter := bson.M{"colName": "users"}
	// find all returns a cursor
	cursor, err := Collection.Find(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M
	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(cursor, context.Background())
	return movies
}
