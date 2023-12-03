include .env.local
export

run:
	go run ./cmd/api/main.go

migrate-up:
	@docker compose --profile tools run --rm migrate up $(step)

migrate-down:
	@docker compose --profile tools run --rm migrate down $(step)
