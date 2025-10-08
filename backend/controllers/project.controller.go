package controllers

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"

	"main/models"
)

func CreateProject(client *mongo.Client, project models.Project) (string, error) {
	collection := client.Database("gdp-nexus").Collection("projects")

	newProject, err := collection.InsertOne(context.TODO(), project)
	if err != nil {
		return "", fmt.Errorf("error inserting project: %w", err)
	}

	return newProject.InsertedID.(string), nil
}
