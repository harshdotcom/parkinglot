package service

import (
	"fmt"
	"harshdotcom/parkinglot/models"
	"harshdotcom/parkinglot/storage"
)

func CreateUser(user models.User) models.User {
	user.ID = storage.UserIdCounter
	storage.UserIdCounter++
	storage.Users = append(storage.Users, user)
	fmt.Println(storage.Users)
	return user
}
