package usecase

import (
	"errors"
	domain "taskmanagement/Domain"
	repository "taskmanagement/Repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllTask(userId primitive.ObjectID) ([]domain.Task, error) {

	tasks, err := repository.GetTasks(userId)
	if err != nil {
		return []domain.Task{}, err
	}

	return tasks, nil
}

func GetTaskByID(taskId primitive.ObjectID, userId primitive.ObjectID) (domain.Task, error) {

	task, err := repository.GetTaskByID(taskId, userId)
	if err != nil {
		return domain.Task{}, err
	}

	return task, nil
}

func CreateTask(task domain.Task, userId primitive.ObjectID) (domain.Task, error) {
	createdTask, ok := repository.CreateTask(task, userId)
	if !ok {
		return domain.Task{}, errors.New("could not create task")
	}
	return createdTask, nil
}

func UpdateTask(task domain.Task, userId primitive.ObjectID) (domain.Task, error) {
	updatedTask, ok := repository.UpdateTask(task, userId)
	if !ok {
		return domain.Task{}, errors.New("could not update task")
	}
	return updatedTask, nil
}

func DeleteTask(taskId primitive.ObjectID, userId primitive.ObjectID) error {
	ok := repository.DeleteTask(taskId, userId)
	if !ok {
		return errors.New("could not delete task")
	}
	return nil
}
