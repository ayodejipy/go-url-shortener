ifneq (,$(wildcard ./.env))
	include .env
	export
endif

# alias air='~/go/bin/air'

build:
	go build -o bin/api cmd/api/main.go && bin/api

run:
	go run cmd/api/main.go

up:
	docker-compose up -d 

down:
	docker-compose down

createdb:
	docker exec -it go-api-dev-db-1 createdb --username=postgres url-shortener

migratecreate:
	migrate create -ext sql -dir $(MIGRATION_PATH) -seq $(name)

migrateup:
	migrate -path $(MIGRATION_PATH) -database $(DATABASE_PSQL_URL) -verbose up

migratedown:
	migrate -path $(MIGRATION_PATH) -database $(DATABASE_PSQL_URL) -verbose down

sqlc:
	sqlc generate

.PHONY: build run up down migratecreate migrateup migratedown sqlc