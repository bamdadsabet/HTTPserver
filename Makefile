include .env

export PGPASSWORD=$(DB_PASSWORD)

run: 
	@go run main.go

buld:
	@go build main.go

migrate-up:
	@echo "Creating tables..."
	@psql -h $(DB_HOST) -U $(DB_USER) -d $(DB_NAME) -f db/query/migrate-up.sql
	@echo "Unsetting PGPASSWORD..."
	@unset PGPASSWORD

migrate-down:
	@echo "Creating tables..."
	@psql -h $(DB_HOST) -U $(DB_USER) -d $(DB_NAME) -f db/query/migrate-down.sql
	@echo "Unsetting PGPASSWORD..."
	@unset PGPASSWORD