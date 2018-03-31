package main

import (
	r "github.com/maikeulb/city-bikes/redis"
	"os"
)

func main() {
	a := App{}
	a.InitializeServer()
	r.InitializeRedis(os.Getenv("REDIS_HOST"))
	a.Run(":5000")
}
