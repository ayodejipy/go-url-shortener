package main

import (
	"fmt"
	"rest/api/internals/cache"
	"rest/api/internals/config"
)


func main() {
	config, _ := config.LoadConfig()

	cache.Init(config)
	err := cache.Set("Name", "superb important value!")
	fmt.Println("Cache error: ", err)
	
	val, err := cache.Get("Name")
	if err != nil {
		fmt.Println("Get Cache err: ", err)
	}
	fmt.Println("Cache returned val: ", val)

	fmt.Println("Entry point to our api")
}