package repository

import (
	"context"
	"errors"
	domain "taskmanagement/Domain"
	infrastructure "taskmanagement/Infrastructure"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func getTaskCollection() (*mongo.Collection, error) {
	if infrastructure.Client == nil {
		return nil, errors.New("mongodb not initialized")
	}
	return infrastructure.Client.Database(infrastructure.DBName).Collection("task"), nil
}

func GetTasks(userID primitive.ObjectID) ([]domain.Task, error) {

	collection, err := getTaskCollection()
	if err != nil {
		return nil, err
	}

	filter := bson.M{"user_id": userID}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	var tasks []domain.Task
	if err := cursor.All(context.Background(), &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func GetTaskByID(taskID primitive.ObjectID, userID primitive.ObjectID) (domain.Task, error) {

	collection, err := getTaskCollection()
	if err != nil {
		return domain.Task{}, err
	}

	var task domain.Task
	err = collection.FindOne(context.Background(), bson.M{"_id": taskID, "user_id": userID}).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.Task{}, errors.New("Task not found")
		}

		return domain.Task{}, errors.New("Server error")

	}

	return task, nil
}

func CreateTask(task domain.Task, userID primitive.ObjectID) (domain.Task, bool) {

	collection, err := getTaskCollection()
	if err != nil {
		return domain.Task{}, false
	}

	task.UserId = userID
	_, err = collection.InsertOne(context.Background(), task)
	if err != nil {
		return domain.Task{}, false
	}

	return task, true
}

func UpdateTask(task domain.Task, userID primitive.ObjectID) (domain.Task, bool) {

	collection, err := getTaskCollection()
	if err != nil {
		return domain.Task{}, false
	}

	filter := bson.M{"_id": task.ID, "user_id": userID}
	err = collection.FindOneAndUpdate(context.Background(), filter, bson.M{"$set": task}).Decode(&task)
	if err != nil {
		return domain.Task{}, false
	}

	return task, true
}

func DeleteTask(taskID primitive.ObjectID, userID primitive.ObjectID) bool {

	collection, err := getTaskCollection()
	if err != nil {
		return false
	}

	filter := bson.M{"_id": taskID, "user_id": userID}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil || result.DeletedCount == 0 {
		return false
	}

	return true
}
