package main

import (
	"github.com/maikeulb/city-bike/redis"
)

func main() {
	a := App{}
	a.Initialize()
	a.Run(":5000")
	redis.SetupRedisCodec()
}
