MIGRATE_CMD=go run migrations/migrate/migrate.go

.PHONY: migrate-up migrate-down migrate-reset

migrate-up:
	$(MIGRATE_CMD) -direction=up

migrate-down:
	$(MIGRATE_CMD) -direction=down

migrate-reset:
	$(MIGRATE_CMD) -direction=reset