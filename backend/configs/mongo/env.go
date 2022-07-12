package mongo

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const DB_NAME = "chat"  // the chat app database name

/**
 * Used to read MongoURI from .env file, this URI is used to connect to MongoDB database.
 */
var GetEnvMongoUri = func() (*string, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
		fmt.Println("> Error loading .env file")

		return nil, err
	}

	uri := os.Getenv("MONGOURI")
	return &uri, nil
}