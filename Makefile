build:
	go build -o bin/api cmd/api/main.go && bin/api

run:
	go run cmd/api/main.go

up:
	docker-compose up -d 

down:
	docker-compose down


.PHONY: build run up down