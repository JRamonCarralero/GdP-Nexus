package controllers

import (
	"context"
	"fmt"
	"main/models"
	"main/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreateNumberTaskIssue(client *mongo.Client, projectId string) (string, error) {
	collection := client.Database("gdp-nexus").Collection("nextTaskNumber")

	pid, err := utils.StringAObjectID(projectId)
	if err != nil {
		return "", err
	}

	newNumber, err := collection.InsertOne(context.TODO(), models.NextTaskNumber{Project: pid, Number: 1})
	if err != nil {
		return "", fmt.Errorf("error inserting project: %w", err)
	}

	return newNumber.InsertedID.(string), nil
}
