package helpers

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DbName = "ShowMovies"
)

var clientInstance *mongo.Client
var clientInstanceError error
var mongoOnce sync.Once

func Setup() (*mongo.Client, error) {
	mongoOnce.Do(func() {

		err := godotenv.Load(".env")

		if err != nil {
			println("could not find the env file")
		}
		// get the url from the .env
		uri := os.Getenv("MONGO_URI")
		if uri == "" {
			log.Fatal("provide a mongo uri variable")
		}
		// apply the url
		clientOptions := options.Client().ApplyURI(uri)
		// connect to mongodb
		client, er := mongo.Connect(context.TODO(), clientOptions)
		// check for err
		if er != nil {
			clientInstanceError = er
		}
		// Check the connection
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = client
	})
	return clientInstance, clientInstanceError

}
