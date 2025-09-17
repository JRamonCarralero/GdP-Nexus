package controllers

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"main/models"
)

func CreateUser(client *mongo.Client, user models.User) error {
	collection := client.Database("gdp-nexus").Collection("users")

	var existingUser models.User
	filter := bson.M{"email": user.Email}
	err := collection.FindOne(context.TODO(), filter).Decode(&existingUser)
	if err == nil {
		return fmt.Errorf("el email ya está registrado")
	}

	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Printf("Error al insertar el usuario: %v", err)
		return fmt.Errorf("error al guardar el usuario en la base de datos")
	}
	return nil
}

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
