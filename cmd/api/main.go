package main

import (
	"rest/api/internals/api"
	"rest/api/internals/config"
)


func main() {
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	// init server
	server := api.NewServer(config)

	server.Start(config.AppPort)
}