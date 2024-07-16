ifneq (,$(wildcard ./.env))
    include .env
    export
endif

DB_M_NAME ?= init
GOOSE_DRIVER=postgres
GOOSE_DBSTRING="host=${DB_HOST} user=${DB_USER} password=${DB_PASS} dbname=${DB_NAME} port=${DB_PORT} sslmode=disable"

run:
	@go run src/*.go

test:
	@go test ./... -v

add-migration:
	@goose -dir ${MIGRATION_PATH} create ${DB_M_NAME} sql

migrate-up:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=${GOOSE_DBSTRING} goose -dir ${MIGRATION_PATH} up
