package db

import (
	"context"
	"os"
	"time"
	"user-service/src/helpers"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var databaseName string

// DBinstance func
func DBinstance() *mongo.Client {
	log := helpers.GetLogger()

	mongoDb, err := helpers.GetEnvStringVal("DATABASE_URL")
	if err != nil {
		log.Error("Failed to load environment variable : DATABASE_URL")
		println("Failed to load environment variable : DATABASE_URL " + err.Error())
		log.Debug(err.Error())
		os.Exit(1)
	}
	databaseName, err = helpers.GetEnvStringVal("DATABASE_NAME")
	if err != nil {
		log.Error("Failed to load environment variable : DATABASE_NAME")
		println("Failed to load environment variable : DATABASE_NAME " + err.Error())
		log.Debug(err.Error())
		os.Exit(1)
	}

	log.Info("Connecting to Database via URL : " + mongoDb)
	println("Connecting to Database via URL : " + mongoDb)
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoDb))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		print("Failed to connect to mongoDB!!!")
		log.Fatal(err)
		println(err)
	}
	log.Info("Connected to mongoDB!!!")
	println("Connected to mongoDB!!!")

	return client
}

func GetClientConnection() *mongo.Client {
	return DBinstance()
}

// OpenCollection is a  function makes a connection with a collection in the database
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {

	var collection *mongo.Collection = client.Database(databaseName).Collection(collectionName)

	return collection
}
