package service

import (
	"awesomeProject/db"
	model "awesomeProject/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Page struct {
	Limit      int64  `json:"limit",bson:"limit"`
	LastId     string `json:"last_id",bson:"last_id"`
	Asc        bool   `json:"asc",bson:"asc"`
	Collection string
	Model      interface{}
}

func GetAll(page Page) ([]bson.M, error) {
	database, ctx, client := db.GetDatabase()
	defer client.Disconnect(ctx)
	usersCollection := database.Collection(page.Collection)
	//filter := bson.D{{"_id", object}}
	var userResult []bson.M
	findOptions := options.Find()
	findOptions.SetLimit(page.Limit)
	//findOptions.SetSort(bson.D{{"duration", -1}})
	sortCursor, err := usersCollection.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	if err = sortCursor.All(ctx, &userResult); err != nil {
		log.Fatal(err)
	}
	return userResult, err
}

func GetOne(object *interface{}, col string) (bson.M, error) {
	database, ctx, client := db.GetDatabase()
	defer client.Disconnect(ctx)
	usersCollection := database.Collection(col)
	var userResult bson.M
	var err error
	original, ok := (*object).(model.User)
	if ok {
		filter := bson.D{{"_id", bson.M{"$eq": original.Id}}}
		err = usersCollection.FindOne(ctx, filter).Decode(&userResult)
	} else {
		err = usersCollection.FindOne(ctx, object).Decode(&userResult)
	}
	if err != nil {
		log.Fatal(err)
	}
	return userResult, err
}

func DeleteOne(object interface{}, col string) (*mongo.DeleteResult, error) {
	database, ctx, client := db.GetDatabase()
	defer client.Disconnect(ctx)
	usersCollection := database.Collection(col)
	result, err := usersCollection.DeleteOne(ctx, object)
	if err != nil {
		log.Fatal(err)
	}
	return result, err
}
func UpdateOne(object interface{}, newObject interface{}, col string) (*mongo.UpdateResult, error) {
	database, ctx, client := db.GetDatabase()
	defer client.Disconnect(ctx)
	usersCollection := database.Collection(col)

	result, err := usersCollection.UpdateOne(
		ctx,
		object,
		bson.D{
			{"$set", newObject},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	return result, err
}
func CreateOne(object interface{}, col string, args ...Page) (*mongo.InsertOneResult, error) {
	database, ctx, client := db.GetDatabase()
	defer client.Disconnect(ctx)
	usersCollection := database.Collection(col)

	//userResult, err := usersCollection.InsertOne(ctx, seedUser())
	userResult, err := usersCollection.InsertOne(ctx, object)
	if err != nil {
		log.Fatal(err)
	}
	return userResult, err
}
