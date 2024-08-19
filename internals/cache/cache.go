package cache

import (
	"context"
	"fmt"
	"rest/api/internals/config"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client
var ctx = context.Background()

func Init(config *config.Config) {
	address := fmt.Sprintf("%s:%v", config.Redis_Host, config.Redis_Port)

	client := redis.NewClient(&redis.Options{
		Addr: address,
		Password: "",
		DB: 0, // use default db
	})

	ping, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Cache error: ", err.Error())
		return
	}

	fmt.Println("Cache test: ", ping)
}

func Set(key string, value string) error {
	err := client.Set(ctx, key, value, 0).Err()
	if err != nil {
		fmt.Printf("Failed to set value in the redis instance: %v \n", err.Error())
		return err
	}
	return nil
}

func Get(key string) (string, error) {
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}