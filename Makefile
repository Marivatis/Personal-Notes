# Load .env.local, ignoring comments and empty lines
ifneq (,$(wildcard .env.local))
	include .env.local
	export $(shell grep -v '^#' .env.local | cut -d= -f1)
endif

MIGRATIONS_PATH := ./migrations

run:
	go run cmd/main.go

migrate-up:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" up

migrate-down:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_URL)" down
