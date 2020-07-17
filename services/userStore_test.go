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
	oid, _ := primitive.ObjectIDFromHex("5f1116fd14137782732387de")
	user, _ := GetUser(model.User{Id: oid})
	if user.Id != oid {
		t.Errorf("Expect %+v, but got %+v\n", user.Id, oid)
	}
}
func TestUpdateUser(t *testing.T) {
	//oid, _ := primitive.ObjectIDFromHex("5f1116fd14137782732387de")
	//newUser,_ := model.CreateObject(model.User{}).(model.User)
	//result, _ := UpdateUser(model.User{Id: oid}, newUser)
	//
	////fmt.Printf("%+v\n", user)
	//if result.ModifiedCount < 1 {
	//	t.Errorf("%d document modifed!\n", result.ModifiedCount)
	//}
}
func TestGetAll(t *testing.T) {
}
