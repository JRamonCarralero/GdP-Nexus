package controllers

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"main/models"
	"main/types"
	"main/utils"
)

func CreateProject(client *mongo.Client, project types.ProjectRequest) (string, error) {
	collection := client.Database("gdp-nexus").Collection("projects")

	newProject, err := collection.InsertOne(context.TODO(), project)
	if err != nil {
		return "", fmt.Errorf("error inserting project: %w", err)
	}

	return newProject.InsertedID.(string), nil
}

func GetProjects(client *mongo.Client) ([]models.Project, error) {
	collection := client.Database("gdp-nexus").Collection("projects")

	var projects []models.Project
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("error finding projects: %w", err)
	}

	err = cursor.All(context.TODO(), &projects)
	if err != nil {
		return nil, fmt.Errorf("error decoding projects: %w", err)
	}

	return projects, nil
}

func GetProjectById(client *mongo.Client, id string) (models.Project, error) {
	collection := client.Database("gdp-nexus").Collection("projects")

	var project models.Project
	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&project)
	if err != nil {
		return models.Project{}, fmt.Errorf("error finding project: %w", err)
	}

	return project, nil
}

func UpdateProject(client *mongo.Client, id string, project types.ProjectUpdateRequest) error {
	collection := client.Database("gdp-nexus").Collection("projects")

	pid, err := utils.StringAObjectID(id)
	if err != nil {
		return err
	}

	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": pid}, bson.M{"$set": project})
	if err != nil {
		return fmt.Errorf("error updating project: %w", err)
	}

	return nil
}

func DeleteProject(client *mongo.Client, id string) error {
	collection := client.Database("gdp-nexus").Collection("projects")

	pid, err := utils.StringAObjectID(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": pid})
	if err != nil {
		return fmt.Errorf("error deleting project: %w", err)
	}

	return nil
}
