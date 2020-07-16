package model

import (
	"fmt"
	"github.com/bxcodec/faker"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	Id             primitive.ObjectID `faker:"-" json:"_id,omitempty" bson:"_id,omitempty"`
	Name           string             `faker:"name" json:"name" bson:"name"`
	Password       string             `faker:"password" json:"password" bson:"password"`
	Phone          string             `faker:"phone_number" json:"phone" bson:"phone"`
	Email          string             `faker:"email" json:"email" bson:"email"`
	EmailConfirmed bool               `json:"emailConfirmed" bson:"emailConfirmed"`
	Token          string             `faker:"jwt" json:"token" bson:"token"`
	CreatedAt      time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt      time.Time          `faker:"-" json:"updatedAt" bson:"updatedAt"`
	DeletedAt      time.Time          `faker:"-" json:"deletedAt" bson:"deletedAt"`
}
type Guest struct {
	UserId    primitive.ObjectID `faker:"-" json:"user_id" bson:"userId"`
	Phone     string             `faker:"phone_number" json:"phone" bson:"phone"`
	Confirmed bool               `json:"confirmed" bson:"confirmed"`
}

type Event struct {
	Id        primitive.ObjectID   `faker:"-" json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string               `json:"name" bson:"name"`
	IsVirtual bool                 `json:"is_virtual" bson:"isVirtual"`
	IsPrivate bool                 `json:"is_private" bson:"isPrivate"`
	Size      int                  `json:"size" bson:"size"`
	Owner     primitive.ObjectID   `json:"owner_id,omitempty" bson:"owner_id,omitempty"`
	Guests    []primitive.ObjectID `json:"guests" bson:"guests"`
	Intro     string               `faker:"sentence" json:"intro" bson:"intro"`
	CreatedAt time.Time            `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time            `faker:"-" json:"updatedAt" bson:"updatedAt"`
	DeletedAt time.Time            `faker:"-" json:"deletedAt" bson:"deletedAt"`
	StartAt   time.Time            `json:"startAt" bson:"startAt"`
	EndAt     time.Time            `json:"endAt" bson:"endAt"`
	Locations []Address            `json:"locations" bson:"locations"`
	Chats     []Message            `json:"chats"`
}
type Message struct {
	UserId  primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Name    string             `faker:"name" json:"name" bson:"name"`
	Message string             `faker:"sentence" json:"message" bson:"message"`
}

type Address struct {
	Street string `json:"street" bson:"street"`
	City   string `json:"city" bson:"city"`
	State  string `json:"state" bson:"state"`
}
type Zoom struct {
	ZoomId   string `json:"zoom_id" bson:"zoom_id"`
	Password string `json:"password" bson:"password"`
}

func CreateObjects(a interface{}, num int) []interface{} {
	v := make([]interface{}, num)
	for i := 0; i < num; i++ {
		v[i] = CreateObject(a)
	}
	return v
}

func CreateObject(a interface{}) interface{} {
	err := faker.FakeData(&a)
	if err != nil {
		fmt.Println(err)
	}
	return &a
}
