package main

import (
	"awesomeProject/db"
	model "awesomeProject/models"
	"fmt"
	"log"
)

func main() {
	database, ctx, client := db.GetDatabase()
	defer client.Disconnect(ctx)
	usersCollection := database.Collection("users")
	eventsCollection := database.Collection("events")
	if err := usersCollection.Drop(ctx); err != nil {
		log.Fatal(err)
	}

	if err := eventsCollection.Drop(ctx); err != nil {
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
