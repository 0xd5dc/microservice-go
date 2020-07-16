package service

import (
	model "awesomeProject/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestGetUser(t *testing.T) {
	e := model.CreateObject(model.User{})
	result, _ := CreateOne(e, "users")
	user, _ := GetUser(&e)
	if user.Id != result.InsertedID {
		t.Errorf("Expect %+v, but got %+v\n", user.Id, result.InsertedID)
	}
}
func TestGetUserByID(t *testing.T) {

	oid, _ := primitive.ObjectIDFromHex("5f10c8e7f1a937281e8c9ac7")
	user, _ := GetUser(model.User{Id: oid})
	if user.Id != oid {
		t.Errorf("Expect %+v, but got %+v\n", user.Id, oid)
	}
}
