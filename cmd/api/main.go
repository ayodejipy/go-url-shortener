package main

import (
	"fmt"
	"rest/api/internals/api"
	"rest/api/internals/config"
	"rest/api/internals/utils"
)


func main() {
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	a := utils.Auth{}
	val, err := a.GenerateRandomCode(32)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Generated code: %v: ", val)

	// init server
	server := api.NewServer(config)

	server.Start(config.AppPort)
}