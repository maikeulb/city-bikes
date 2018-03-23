package main

import (
	"github.com/maikeulb/city-bike/redis"
)

func main() {
	a := App{}
	a.Initialize()
	redis.SetupRedisCodec()
	a.Run(":5000")
}
