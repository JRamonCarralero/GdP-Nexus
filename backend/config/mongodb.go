package config

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

// ConnectDB connects to a MongoDB database and pings it to ensure the connection is successful.
// The function takes no parameters and returns an error if something goes wrong.
// If the connection to the database is successful, a nil error is returned.
// The function stores the MongoDB client in the global DB variable.
//
// @return error
func ConnectDB() error {
	clientOptions := options.Client().ApplyURI(os.Getenv("DB_URI"))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return fmt.Errorf("Error connecting to MongoDB: %w", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return fmt.Errorf("Error pinging MongoDB: %w", err)
	}

	fmt.Println("MongoDB connected!")

	DB = client
	return nil
}
