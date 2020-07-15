package main

import (
	"fmt"
	"github.com/bxcodec/faker"
)

type User struct {
	Name           string `faker:"name",json:"name"`
	Password       string `faker:"password",json:"password"`
	Phone          string `faker:"phone_number",json:"phone"`
	Email          string `faker:"email",json:"email"`
	EmailConfirmed bool   `json:"emailConfirmed"`
	Token          string `faker:"jwt",json:"token"`
	CreatedAt      string `faker:"timestamp",json:"createdAt"`
	UpdatedAt      string `faker:"timestamp",json:"updatedAt"`
}
type Guest struct {
	UserId    string `json:"user_id",bson:"userId"`
	Phone     string `faker:"phone_number",json:"phone",bson:"phone"`
	Confirmed bool   `json:"confirmed",bson:"confirmed"`
}

type Event struct {
	Name      string    `json:"name",bson:"name"`
	IsVirtual bool      `json:"is_virtual",bson:"isVirtual"`
	IsPrivate bool      `json:"is_private",bson:"isPrivate"`
	Size      int       `json:"size",bson:"size"`
	Owner     string    `json:"owner",bson:"owner"`
	Guests    []Guest   `json:"guests",bson:"guests"`
	Intro     string    `faker:"sentence",json:"intro",bson:"intro"`
	CreatedAt string    `faker:"timestamp",json:"created_at",bson:"createdAt"`
	UpdatedAt string    `faker:"timestamp",json:"updated_at",bson:"updatedAt"`
	DeletedAt string    `faker:"timestamp",json:"deleted_at",bson:"deletedAt"`
	StartAt   string    `faker:"timestamp",json:"start_at",bson:"startAt"`
	EndAt     string    `faker:"timestamp",json:"end_at",bson:"endAt"`
	Locations []Address `json:"locations",bson:"locations"`
	Chats     []Message `json:"chats"`
}
type Message struct {
	UserId  string `json:"user_id",bson:"userId"`
	Name    string `faker:"name",json:"name",bson:"name"`
	Message string `faker:"sentence",json:"message",bson:"message"`
}

type Address struct {
	Street string `json:"street",bson:"street"`
	City   string `json:"city",bson:"city"`
	State  string `json:"state",bson:"state"`
}
type Zoom struct {
	ZoomId   string `json:"zoom_id",bson:"zoomId"`
	Password string `json:"password",bson:"password"`
}

func main() {
	v := createObjects(Zoom{})
	fmt.Print(v)
}

func createObjects(a interface{}) []interface{} {
	v := make([]interface{}, 100)
	for i := 0; i < 100; i++ {
		v[i], _ = createObject(a)
	}
	return v
}

func createObject(a interface{}) (interface{}, error) {
	err := faker.FakeData(&a)
	if err != nil {
		fmt.Println(err)
	}
	return &a, err
}
