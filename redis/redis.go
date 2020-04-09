package redis

import (
	"fmt"

	redis "github.com/go-redis/redis/v7"
)

// RClient create client
func RClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	return client
}

// Ping db
func Ping(client *redis.Client) error {
	pong, err := client.Ping().Result()
	if err != nil {
		return err
	}
	fmt.Println(pong, err)
	// Output: PONG <nil>

	return nil
}

// Set key, value
func Set(client *redis.Client, key string, value string) error {
	err := client.Set(key, value, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

// Get value
func Get(client *redis.Client, key string) (string, error) {
	value, err := client.Get(key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}
