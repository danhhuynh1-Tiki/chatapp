package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	mongoEnv "jwt-demo/configs/mongo"
	config "jwt-demo/configs/mongo"
)

/**
 * Connect to the MongoDB database.
 */
var ConnectDatabase = func() *mongo.Client {
	uri, err := mongoEnv.GetEnvMongoUri()
	
	if err != nil {
		log.Fatal(err)
		fmt.Println("> Something went wronng in configurations to MongoDB database...")

		return nil
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(*uri))

	// Timeout is 10 seconds when we try to connect to MongoDB
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err = client.Connect(ctx); err != nil { // connect to MongoDB failed
		log.Fatal(err)
		fmt.Println("> Can not connect to MongoDB database...")

		return nil
	}

	// Ping to MongoDB to test connection
	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
		fmt.Println("> Can not ping to MongoDB database...")

		return nil
	}

	fmt.Println("> Connected to MongoDB successfully.")
	return client
}

/**
 * Used to refer to a specific collection in MongoDB
 */
var GetCollection = func(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database(config.DB_NAME).Collection(collectionName)
}