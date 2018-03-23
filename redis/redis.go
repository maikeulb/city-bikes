package redis

import (
	"encoding/json"

	"github.com/go-redis/cache"
	"github.com/go-redis/redis"
)

var Codec *cache.Codec

func SetupRedisCodec() {
	client := redis.NewClient(&redis.Options{
		Addr:     "172.17.0.3:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
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
