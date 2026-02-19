package repository

import (
	"context"
	"errors"
	domain "taskmanagement/Domain"
	infrastructure "taskmanagement/Infrastructure"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func getUserCollection() (*mongo.Collection, error) {
	if infrastructure.Client == nil {
		return nil, errors.New("mongodb not initialized")
	}
	return infrastructure.Client.Database(infrastructure.DBName).Collection("user"), nil
}

func GetUserByEmail(email string) (domain.User, error) {

	collection, err := getUserCollection()
	if err != nil {
		return domain.User{}, err
	}

	var user domain.User

	err = collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil

}

func CreateUser(user domain.User) error {
	collection, err := getUserCollection()
	if err != nil {
		return err
	}

	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}
