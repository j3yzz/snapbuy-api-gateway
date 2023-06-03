package cache

import "github.com/go-redis/redis"

type Handler struct {
	cache *redis.Client
}

func Init(addr string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	return client
}
