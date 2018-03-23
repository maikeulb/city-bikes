package redis

import (
	"encoding/json"

	"github.com/go-redis/cache"
	"github.com/go-redis/redis"
)

var Codec *cache.Codec

func SetupRedisCodec() {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": "172.17.0.3:6379",
		},
	})

	Codec = &cache.Codec{
		Redis: ring,

		Marshal: func(v interface{}) ([]byte, error) {
			return json.Marshal(v)
		},
		Unmarshal: func(b []byte, v interface{}) error {
			return json.Unmarshal(b, v)
		},
	}
}
