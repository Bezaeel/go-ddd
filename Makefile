ifneq (,$(wildcard ./.env))
    include .env
    export
endif

DB_M_NAME ?= init
GOOSE_DRIVER=postgres
GOOSE_DBSTRING="host=${POSTGRES_HOST} user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DATABASE} port=${POSTGRES_PORT} sslmode=disable"

run:
	@go run src/*.go

test:
	@go test ./... -v

add-migration:
	@goose -dir ${MIGRATION_PATH} create ${DB_M_NAME} sql

migrate-up:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=${GOOSE_DBSTRING} goose -dir ${MIGRATION_PATH} up
