include .env

MIGRATION_PATH = ./cmd/migrate/migrations

.PHONY: migrate-create
migrate-create:
	@migrate create -seq -ext sql -dir $(MIGRATION_PATH) $(name)

.PHONY: migrate-up 
migrate-up: 
	@migrate -path $(MIGRATION_PATH) -database "$(DB_ADDR)" up

.PHONY: migrate-down
migrate-down: 
	@migrate -path $(MIGRATION_PATH) -database "$(DB_ADDR)" down $(steps)
