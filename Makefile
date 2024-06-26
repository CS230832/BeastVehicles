build:
	@go build -o bin/app.exe cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/app.exe

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down

up:
	@docker compose up -d
down:
	@docker compose down