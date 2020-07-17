package service

import (
	model "awesomeProject/models"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// repo set to the collection of users

func All() error {
	// use drive to get list of users in db and return users by part and error
	// repo.getAll(data: user)
	return errors.New("Not implemented")
}
func GetUser(user interface{}) (model.User, error) {
	// use drive to get data of a user in db and return user and error
	user_bson, err := GetOne(&user, "users")
	model := bsonToModel(user_bson)
	return model, err
}

func bsonToModel(user_bson bson.M) model.User {
	var s model.User
	bsonBytes, _ := bson.Marshal(user_bson)
	bson.Unmarshal(bsonBytes, &s)
	return s
}

func UpdateUser(filter bson.D, update bson.D) (*mongo.UpdateResult, error) {
	// use drive to update updatedAt for a user in db and return user and error
	// repo.updateOne(data: user)
	result, err := UpdateOne(filter, update, "users")
	return result, err
}

func DeleteUser() error {
	// use drive to update deleteAt for a user in db and return deleteAt and error
	// repo.deleteOne(data: user)
	return errors.New("Not implemented")
}

func CreateUser(user interface{}) (primitive.ObjectID, error) {
	// use drive to create a user in db and return user id and error if any
	result, err := CreateOne(&user, "users")
	insertedId, _ := result.InsertedID.(primitive.ObjectID)
	return insertedId, err
}
