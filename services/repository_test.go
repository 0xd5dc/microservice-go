package service

import (
	model "awesomeProject/models"
	"testing"
)

func TestRepository(t *testing.T) {
	t.Run("create_a user", createTest(model.User{}, "users"))
	t.Run("delete_a_user", deleteTest(model.User{}, "users"))
	t.Run("create_a_event", createTest(model.Event{}, "events"))
	t.Run("delete_a_event", deleteTest(model.Event{}, "events"))
	t.Run("update_a_user", updateTest(model.User{}, "users"))
	t.Run("find_a_user", findTest(model.User{}, "users"))
	t.Run("find_a_event", findTest(model.Event{}, "events"))
	t.Run("update_a_event", updateTest(model.Event{}, "events"))
}

func findTest(object interface{}, col string) func(t *testing.T) {
	return func(t *testing.T) {
		obj := model.CreateObject(object)
		CreateOne(obj, col)
		result, _ := GetOne(obj, col)
		if result == nil {
			t.Errorf("0 object found!")
		}
	}
}

func deleteTest(object interface{}, col string) func(t *testing.T) {
	return func(t *testing.T) {
		obj := model.CreateObject(object)
		CreateOne(obj, col)
		result, _ := DeleteOne(obj, col)
		if result.DeletedCount < 1 {
			t.Errorf("%d 0 document deleted!", result.DeletedCount)
		}
	}
}

func createTest(object interface{}, col string) func(t *testing.T) {
	return func(t *testing.T) {
		obj := model.CreateObject(object)
		// generated a fake obj
		result, _ := CreateOne(obj, col)
		if result.InsertedID == nil {
			t.Errorf("0 document created!")
		}
	}
}
func updateTest(object interface{}, col string) func(t *testing.T) {
	return func(t *testing.T) {
		obj := model.CreateObject(object)
		CreateOne(obj, col)
		newObj := model.CreateObject(object)
		// generated a fake obj
		result, _ := UpdateOne(obj, newObj, col)
		if result.MatchedCount < 1 {
			t.Errorf("0 document found!")
		}
		if result.ModifiedCount < 1 {
			t.Errorf("O document modified!")
		}
	}
}
