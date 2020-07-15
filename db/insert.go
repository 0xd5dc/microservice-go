package main

import (
	model "awesomeProject/models"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://10.0.3.6:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	database := client.Database("minimalGraphql")
	usersCollection := database.Collection("users")
	eventsCollection := database.Collection("events")
	if err = usersCollection.Drop(ctx); err != nil {
		log.Fatal(err)
	}

	if err = eventsCollection.Drop(ctx); err != nil {
		log.Fatal(err)
	}
	fmt.Print("DB Purged\n")
	users := model.CreateObjects(model.User{}, 1000)

	//userResult, err := usersCollection.InsertOne(ctx, seedUser())
	userResult, err := usersCollection.InsertMany(ctx, users)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d users inserted!\n", len(userResult.InsertedIDs))

	events := model.CreateObjects(model.Event{}, 100)
	eventResult, err := eventsCollection.InsertMany(ctx, events)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d events inserted!\n", len(eventResult.InsertedIDs))
}
