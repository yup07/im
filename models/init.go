package models

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var Mongo = InitMongo()
var RDB = InitRedis()

func InitMongo() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().SetAuth(options.Credential{
		Username: "admin",
		Password: "123456",
	}).ApplyURI("mongodb://47.115.205.241:27017"))
	if err != nil {
		log.Fatal("mongo 链接错误", err)
		return nil
	}
	return client.Database("im")

}
func InitRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "47.115.205.241:6379",
		Password: "123456",
	})

}
