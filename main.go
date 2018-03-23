package main

import (
	r "github.com/maikeulb/city-bikes/redis"
)

func main() {
	a := App{}
	a.InitializeServer()
	r.InitializeRedis()
	a.Run(":5000")
}
