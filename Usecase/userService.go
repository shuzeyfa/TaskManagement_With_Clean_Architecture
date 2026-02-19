package usecase

import (
	"errors"
	domain "taskmanagement/Domain"
	infrastructure "taskmanagement/Infrastructure"
	repository "taskmanagement/Repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(req domain.RegisterRequest) (domain.User, error) {

	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		12,
	)
	if err != nil {
		return domain.User{}, err
	}

	// check if the user already exist
	_, err = repository.GetUserByEmail(req.Email)
	if err == nil {
		return domain.User{}, errors.New("user already exist")
	}

	user := domain.User{
		ID:       primitive.NewObjectID(),
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     "user",
	}

	err = repository.CreateUser(user)

	return user, nil
}

func LoginUser(req domain.LoginRequest) (string, error) {

	// check if the user not registered
	user, err := repository.GetUserByEmail(req.Email)
	if err != nil {
		return "", errors.New("Invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	); err != nil {
		return "", errors.New("invalid credentials")
	}

	tokenString, err := infrastructure.GenerateJWT(user, req)
	if err != nil {
		return "", errors.New(err.Error())
	}

	return tokenString, nil
}
