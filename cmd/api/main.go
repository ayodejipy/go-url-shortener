package main

import (
	"fmt"
	"rest/api/internals/cache"
	"rest/api/internals/config"
)


func main() {
	config, _ := config.LoadConfig()


	cache.Init(config)
	err := cache.Set("key", "superb important value!")
	fmt.Println("Cache error: ", err)

	fmt.Println("Entry point to our api")
}