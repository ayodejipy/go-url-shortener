package main

import (
	"fmt"
	"rest/api/internals/api"
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
	server := api.NewServer(config)

	server.Start(config.AppPort)
}