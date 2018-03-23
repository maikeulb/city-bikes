package redis

import (
	"encoding/json"

	"github.com/go-redis/cache"
	"github.com/go-redis/redis"
)

var Codec *cache.Codec

func InitializeRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     "172.17.0.3:6379",
		Password: "",
		DB:       0,
	})

	Codec = &cache.Codec{
		Redis: client,

		Marshal: func(v interface{}) ([]byte, error) {
			return json.Marshal(v)
		},
		Unmarshal: func(b []byte, v interface{}) error {
			return json.Unmarshal(b, v)
		},
	}
}
