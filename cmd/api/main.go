package main

import (
	"fmt"
	"rest/api/internals/cache"
	"rest/api/internals/config"
)


func main() {
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	// init cache
	cache.Init(config)
	cacheErr := cache.Set("Name", "superb important value!")
	fmt.Println("Cache error: ", cacheErr)

	// init server
	
	val, err := cache.Get("Name")
	if err != nil {
		fmt.Println("Get Cache err: ", err)
	}
	fmt.Println("Cache returned val: ", val)

	fmt.Println("Entry point to our api")
}