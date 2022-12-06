package client

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/tuongnguyen1209/poly-career-back/config"
)

var Client *redis.Client

func init() {

	config := config.GetConfig()

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port),
		Password: config.Redis.Password,
		DB:       0,
	})

	pong, err := client.Ping().Result()
	Client = client
	fmt.Println("connect redis, ", pong, "<> Err: ", err)
}

func SetValue(name string, value interface{}, time time.Duration) error {
	return Client.Set(name, value, time).Err()
}

func GetValue(name string) (string, error) {
	return Client.Get(name).Result()
}
