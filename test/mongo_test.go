package test

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"im/models"
	"testing"
	"time"
)

func TestFindOne(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().SetAuth(options.Credential{
		Username: "admin",
		Password: "123456",
	}).ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatal(err)
	}
	db := client.Database("im")
	ub := new(models.UserBasic)
	err = db.Collection("user_basic").FindOne(context.TODO(), bson.D{}).Decode(ub)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("ub====", ub)
}
func TestFind(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().SetAuth(options.Credential{
		Username: "admin",
		Password: "123456",
	}).ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatal(err)
	}
	db := client.Database("im")
	urs := make([]*models.UserRoom, 0)
	cursor, err := db.Collection("user_room").Find(context.TODO(), bson.D{})
	for cursor.Next(context.Background()) {
		ur := new(models.UserRoom)
		err = cursor.Decode(ur)
		if err != nil {
			t.Fatal(err)
		}
		urs = append(urs, ur)
	}
	for _, v := range urs {
		fmt.Println("UserRoom", v)

	}
}
