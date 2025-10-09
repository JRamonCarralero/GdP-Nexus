package controllers

import (
	"context"
	"fmt"
	"main/models"
	"main/types"
	"main/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateNumberTaskIssue creates a new next task number document in the database.
//
// The function takes a pointer to a mongo client and a project id as parameters and returns the new task number as a string and an error if something goes wrong.
//
// If an error occurs while inserting the next task number, a 500 Internal Server Error is returned.
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

// GetNextTaskNumber gets the next task number for a given project from the database.
//
// The function takes a pointer to a mongo client and a project id as parameters and returns the next task number as an int and an error if something goes wrong.
//
// If an error occurs while getting the next task number, a 500 Internal Server Error is returned.
func GetNextTaskNumber(client *mongo.Client, projectId string) (int, error) {
	collection := client.Database("gdp-nexus").Collection("nextTaskNumber")

	pid, err := utils.StringAObjectID(projectId)
	if err != nil {
		return 0, err
	}

	var result struct {
		Number int `bson:"number"`
	}

	filter := bson.D{{Key: "project", Value: pid}}
	projection := bson.D{
		{Key: "number", Value: 1},
		{Key: "_id", Value: 0},
	}

	opt := options.FindOne().SetProjection(projection)
	err = collection.FindOne(context.TODO(), filter, opt).Decode(&result)
	if err != nil {
		return 0, err
	}

	_, err = collection.UpdateOne(context.TODO(), filter, bson.M{"$inc": bson.M{"number": 1}})
	if err != nil {
		return 0, err
	}

	return result.Number, nil
}

// CreateTask creates a new task in the database.
//
// The function takes a pointer to a mongo client and a models.TaskRequest
// as parameters and returns the new task id as a string and an error if
// something goes wrong.
//
// If an error occurs while getting the next task number, a 500 Internal
// Server Error is returned.
//
// If an error occurs while inserting the task, a 500 Internal Server Error
// is returned.
func CreateTask(client *mongo.Client, task types.TaskRequest) (string, error) {
	collection := client.Database("gdp-nexus").Collection("tasks")

	taskNumber, err := GetNextTaskNumber(client, task.Project)
	if err != nil {
		return "", err
	}
	assignee, err := utils.StringAObjectID(task.Assignee)
	if err != nil {
		return "", err
	}
	createdBy, err := utils.StringAObjectID(task.CreatedBy)
	if err != nil {
		return "", err
	}
	projectId, err := utils.StringAObjectID(task.Project)
	if err != nil {
		return "", err
	}

	newTask := models.Task{
		ID:          primitive.NewObjectID(),
		Name:        task.Name,
		Description: task.Description,
		Project:     projectId,
		Assignee:    assignee,
		Number:      taskNumber,
		CreatedBy:   createdBy,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}

	resultTask, err := collection.InsertOne(context.TODO(), newTask)
	if err != nil {
		return "", fmt.Errorf("error inserting task: %w", err)
	}

	return resultTask.InsertedID.(string), nil
}

// GetTasks gets all tasks for a given project from the database.
//
// The function takes a pointer to a mongo client and a project id as parameters and returns a slice of models.Task and an error if something goes wrong.
//
// If an error occurs while finding the tasks, a 500 Internal Server Error is returned.
//
// If an error occurs while decoding the tasks, a 500 Internal Server Error is returned.
func GetTasks(client *mongo.Client, projectId string) ([]models.Task, error) {
	collection := client.Database("gdp-nexus").Collection("tasks")

	var tasks []models.Task
	cursor, err := collection.Find(context.TODO(), bson.M{"project": projectId})
	if err != nil {
		return nil, fmt.Errorf("error finding tasks: %w", err)
	}

	err = cursor.All(context.TODO(), &tasks)
	if err != nil {
		return nil, fmt.Errorf("error decoding tasks: %w", err)
	}

	return tasks, nil
}

// GetTaskById gets a task from the database by its ID.
//
// The function takes a pointer to a mongo client and a task ID as parameters and returns a models.Task and an error if something goes wrong.
//
// If an error occurs while finding the task, a 500 Internal Server Error is returned.
//
// If an error occurs while decoding the task, a 500 Internal Server Error is returned.
func GetTaskById(client *mongo.Client, id string) (models.Task, error) {
	collection := client.Database("gdp-nexus").Collection("tasks")

	var task models.Task
	tid, err := utils.StringAObjectID(id)
	if err != nil {
		return models.Task{}, err
	}

	err = collection.FindOne(context.TODO(), bson.M{"_id": tid}).Decode(&task)
	if err != nil {
		return models.Task{}, fmt.Errorf("error finding task: %w", err)
	}

	return task, nil
}

// UpdateTask updates a task by ID.
//
// The function takes a pointer to a mongo client, a task ID, and a types.TaskUpdateRequest as parameters and returns an error if something goes wrong.
//
// If an error occurs while updating the task, a 500 Internal Server Error is returned.
func UpdateTask(client *mongo.Client, id string, task types.TaskUpdateRequest) error {
	collection := client.Database("gdp-nexus").Collection("tasks")

	tid, err := utils.StringAObjectID(id)
	if err != nil {
		return err
	}

	updateDoc := bson.M{}

	if task.Name != nil {
		updateDoc["name"] = *task.Name
	}

	if task.Description != nil {
		updateDoc["description"] = *task.Description
	}
	if task.Assignee != nil {
		assigneeStr := *task.Assignee
		assignee, err := utils.StringAObjectID(assigneeStr)
		if err != nil {
			return err
		}

		updateDoc["assignee"] = assignee
	}

	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": tid}, bson.M{"$set": updateDoc})
	if err != nil {
		return fmt.Errorf("error updating task: %w", err)
	}

	return nil
}

// DeleteTask deletes a task by ID.
//
// The function takes a pointer to a mongo client and a task ID as parameters and returns an error if something goes wrong.
//
// If an error occurs while deleting the task, a 500 Internal Server Error is returned.
func DeleteTask(client *mongo.Client, id string) error {
	collection := client.Database("gdp-nexus").Collection("tasks")

	tid, err := utils.StringAObjectID(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": tid})
	if err != nil {
		return fmt.Errorf("error deleting task: %w", err)
	}

	return nil
}
