package helpers

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "ShowMovies"
const colName = "users"

var Collection *mongo.Collection

func init() {
	// load env files
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
	client, err := mongo.Connect(context.TODO(), clientOptions)
	// check for err
	if err != nil {
		log.Fatal(err)
	}
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(client, context.Background())
	Collection = client.Database(dbName).Collection(colName)
	fmt.Println("mongodb connection success")
}
