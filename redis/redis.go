package redis

import (
	"encoding/json"

	"github.com/go-redis/cache"
	"github.com/go-redis/redis"
)

var Codec *cache.Codec

func InitializeRedis(addr string) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
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
