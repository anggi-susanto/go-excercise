include .env

.PHONY: all
all: build
FORCE: ;

BIN_DIR = $(PWD)/bin

.PHONY: build

clean:
	rm -rf bin/*

dependencies:
	go mod download

build: dependencies build-api build-cmd

create-migration:
	goose -dir=build/migrations create $(file) sql

create-seeder:
	goose -dir=build/seeders create $(file) sql

run-migrations:
	set -a && . ./.env && echo $(DB_DSN) && goose -dir=build/migrations mysql $(DB_DSN) up

rollback-migrations:
	set -a && . ./.env && goose -dir=build/migrations mysql $(DB_DSN) down	

run-seeders:
	set -a && . ./.env && goose -dir=build/seeders -table=goose_seeder_version mysql $(DB_DSN) up

rollback-seeders:
	set -a && . ./.env && goose -dir=build/seeders -table=goose_seeder_version mysql $(DB_DSN) down

serve-rest:
	set -a && . ./.env && go run *.go serve:rest

api-docs:
	~/go/bin/swag init -g cmd/serverest/cmd_serve.go

unit-test:
	set -a && . ./.env && go test -race -v -coverprofile=profile.out ./... $(shell echo $(TEST_FLAGS)) && go tool cover -html=cover.out ; rm -f cover.out

coverage:
	@go test -covermode=count -coverprofile=count.out fmt; rm -f count.out

build-mocks:
	@./build/mock.sh