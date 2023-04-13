package test

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

var ctx = context.Background()

func TestSet(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456", // no password set
		DB:       0,        // use default DB
	})

	err := rdb.Set(ctx, "key", "123", time.Second*30).Err()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGet(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456", // no password set
		DB:       0,        // use default DB
	})

	r, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(r)
}
