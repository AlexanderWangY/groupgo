include .env


setup:
	@echo "Installing required dependencies, Go should be installed..."
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.28.0
	go install github.com/pressly/goose/v3/cmd/goose@v3.24.1
	go install github.com/air-verse/air@v1.61.7

	@echo "Installation successful!\n\n"
	@echo "Setting up local DB (ensure you have Docker installed!)"
	
	@if docker ps -a --filter "name=dev-postgres" | grep -q "dev-postgres"; then \
		echo "Container dev-postgres exists. Starting the container..."; \
		make start-local-db; \
	else \
		echo "Container dev-postgres does not exist. Creating it now!"; \
		docker run --name dev-postgres -e POSTGRES_PASSWORD=postgres -d -p 5432:5432 postgres:15; \
	fi
	@sleep 2
	@echo "Applying migrations..."
	@make migrate-up
	@echo "Migrations done.\n\n"
	@sleep 2
	@echo "Generating sqlc types and queries..."
	@make generate
	@echo "Done generating types!"
	@echo "Happy developing!"

	
run:
	go run ./cmd/server/main.go

generate:
	@sqlc generate

migrate-up:
	@goose -dir db/migrations postgres "$(DATABASE_URL)" up

start-local-db:
	docker start dev-postgres

stop-local-db:
	docker stop dev-postgres
