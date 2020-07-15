package service

import (
	model "awesomeProject/models"
	"fmt"
	"log"
	"testing"
)

func TestGetOne(t *testing.T) {
	user := model.CreateObject(model.User{})
	// generated a fake user
	createdUser, err := CreateOne(user, "users")
	if err != nil {
		log.Fatal(err)
	}
	foundUser, err := GetOne(createdUser, "users")
	if createdUser != foundUser {
		fmt.Printf("CreateOne with ID: %+v, but find:%+v\n", createdUser, foundUser)
	}
}
func TestCreateOne(t *testing.T) {
	user := model.CreateObject(model.User{})
	_, err := CreateOne(user, "users")
	if err != nil {
		log.Fatal(err)
	}
}
