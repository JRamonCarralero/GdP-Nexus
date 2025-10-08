package controllers

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"main/models"
)

// CreateUser creates a new user in the database.
//
// The function takes a pointer to a mongo client and a models.User
// as parameters and returns an error if something goes wrong.
//
// If the email is already registered, a 409 Conflict error is returned.
// If an error occurs while inserting the user, a 500 Internal Server Error
// is returned.
//
// @param client *mongo.Client
// @param user models.User
// @return error
func CreateUser(client *mongo.Client, user models.User) error {
	collection := client.Database("gdp-nexus").Collection("users")

	var existingUser models.User
	filter := bson.M{"email": user.Email}
	err := collection.FindOne(context.TODO(), filter).Decode(&existingUser)
	if err == nil {
		return fmt.Errorf("email is already registered")
	}

	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		return fmt.Errorf("error inserting user: %w", err)
	}
	return nil
}

// GetUserByEmail gets a user from the database by its email.
//
// The function takes a pointer to a mongo client and an email string
// as parameters and returns a models.User and an error.
//
// If the email does not exist, a nil models.User and a nil error are returned.
// If an error occurs while getting the user, the error is returned.
//
// @param client *mongo.Client
// @param email string
// @return models.User, error
func GetUserByEmail(client *mongo.Client, email string) (models.User, error) {
	collection := client.Database("gdp-nexus").Collection("users")

	var user models.User
	filter := bson.M{"email": email}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
