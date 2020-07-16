package service

import (
	model "awesomeProject/models"
	"testing"
)

func TestCreateOneGetOne(t *testing.T) {
	t.Run("create_and_find_a user", createTest(model.User{}, "users"))
	t.Run("create_and_find_a_event", createTest(model.Event{}, "events"))
}

func TestDeleteOne(t *testing.T) {
	t.Run("delete_a_user", deleteTest(model.User{}, "users"))
	t.Run("delete_a_event", deleteTest(model.Event{}, "events"))
}

func deleteTest(object interface{}, col string) func(t *testing.T) {
	return func(t *testing.T) {
		obj := model.CreateObject(object)
		CreateOne(obj, col)
		result, _ := DeleteOne(obj, col)
		if result.DeletedCount < 1 {
			t.Errorf("DeleteOne removed %+v document(s)\n", result.DeletedCount)
		}
	}
}

func createTest(object interface{}, col string) func(t *testing.T) {
	return func(t *testing.T) {
		obj := model.CreateObject(object)
		// generated a fake obj
		createdObject, _ := CreateOne(obj, col)
		foundObject, _ := GetOne(createdObject, col)
		if createdObject != foundObject {
			t.Errorf("CreateOne with ID: %+v, but find:%+v\n", createdObject, foundObject)
		}
	}
}
