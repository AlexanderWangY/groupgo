include .env


run:
	go run ./cmd/server/main.go

generate:
	sqlc generate

migrate-up:
	goose -dir db/migrations postgres "$(DATABASE_URL)" up

start-local-db:
	docker start dev-postgres

stop-local-db:
	docker stop dev-postgres
