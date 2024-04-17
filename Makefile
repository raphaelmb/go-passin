include .env

build:
	@go build -o bin/passin ./cmd

run: build
	@./bin/passin

test:
	@go test -v ./...

create_migration:
	@read -p "Enter migration name: " name; \
		migrate create -ext sql -dir internal/database/migrations $$name

migrate_up:
	migrate -path=internal/database/migrations \
		-database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose up

migrate_down:
	migrate -path=internal/database/migrations \
		-database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose down


.PHONY: build run test create_migration migrate_up migrate_down
