package controllers

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"main/models"
	"main/types"
	"main/utils"
)

// CreateProject creates a new project in the database.
//
// The function takes a pointer to a mongo client and a models.ProjectRequest
// as parameters and returns the new project id as a string and an error if
// something goes wrong.
//
// If an error occurs while getting the owner id, a 500 Internal Server
// Error is returned.
//
// If an error occurs while inserting the project, a 500 Internal Server
// Error is returned.
func CreateProject(client *mongo.Client, project types.ProjectRequest) (string, error) {
	collection := client.Database("gdp-nexus").Collection("projects")

	ownerId, err := utils.StringAObjectID(project.Owner)
	if err != nil {
		return "", err
	}

	membersIds := make([]primitive.ObjectID, len(project.Members))
	for i, memberId := range project.Members {
		membersIds[i], err = utils.StringAObjectID(memberId)
		if err != nil {
			return "", err
		}
	}

	newProject := models.Project{
		ID:          primitive.NewObjectID(),
		Name:        project.Name,
		Description: project.Description,
		Owner:       ownerId,
		Members:     membersIds,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		CreatedBy:   ownerId,
	}

	nProject, err := collection.InsertOne(context.TODO(), newProject)
	if err != nil {
		return "", fmt.Errorf("error inserting project: %w", err)
	}

	return nProject.InsertedID.(string), nil
}

// GetProjects gets all projects from the database.
//
// The function takes a pointer to a mongo client as a parameter and returns a slice of models.Project and an error if something goes wrong.
//
// If an error occurs while finding the projects, a 500 Internal Server Error is returned.
//
// If an error occurs while decoding the projects, a 500 Internal Server Error is returned.
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

// GetProjectById gets a project from the database by its ID.
//
// The function takes a pointer to a mongo client and a project ID as parameters and returns a models.Project and an error if something goes wrong.
//
// If an error occurs while finding the project, a 500 Internal Server Error is returned.
//
// If an error occurs while decoding the project, a 500 Internal Server Error is returned.
func GetProjectById(client *mongo.Client, id string) (models.Project, error) {
	collection := client.Database("gdp-nexus").Collection("projects")

	var project models.Project
	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&project)
	if err != nil {
		return models.Project{}, fmt.Errorf("error finding project: %w", err)
	}

	return project, nil
}

// UpdateProject updates a project by ID.
//
// The function takes a pointer to a mongo client, a project ID, and a types.ProjectUpdateRequest as parameters and returns an error if something goes wrong.
//
// If an error occurs while updating the project, a 500 Internal Server Error is returned.
func UpdateProject(client *mongo.Client, id string, project types.ProjectUpdateRequest) error {
	collection := client.Database("gdp-nexus").Collection("projects")

	pid, err := utils.StringAObjectID(id)
	if err != nil {
		return err
	}

	updateDoc := bson.M{}

	if project.Name != nil {
		updateDoc["name"] = *project.Name
	}
	if project.Description != nil {
		updateDoc["description"] = *project.Description
	}
	if project.Owner != nil {
		ownerStr := *project.Owner
		owner, err := utils.StringAObjectID(ownerStr)
		if err != nil {
			return err
		}

		updateDoc["owner"] = owner
	}
	if project.Members != nil {
		membersIds := make([]primitive.ObjectID, len(*project.Members))
		for i, memberId := range *project.Members {
			membersIds[i], err = utils.StringAObjectID(memberId)
			if err != nil {
				return err
			}
		}
		updateDoc["members"] = membersIds
	}

	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": pid}, bson.M{"$set": updateDoc})
	if err != nil {
		return fmt.Errorf("error updating project: %w", err)
	}

	return nil
}

// DeleteProject deletes a project by ID.
//
// The function takes a pointer to a mongo client and a project ID as parameters and returns an error if something goes wrong.
//
// If an error occurs while deleting the project, a 500 Internal Server Error is returned.
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
