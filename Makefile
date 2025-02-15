include .env
export $(shell sed 's/=.*//' .env)

db-up:
	@docker-compose -f docker-compose.yaml up -d postgres
	@echo "🚀 Database is up and running!"

db-down:
	@docker-compose -f docker-compose.yaml down --volumes
	@echo "💥 Database is down!"

db-drop: db-down
	@docker-compose -f docker-compose.yaml rm -v postgres
	@echo "💀 Database deleted!"

build:
	@templ generate
	@go mod tidy
	@go build -o ./bin/admin-panel ./cmd/api/main.go

run: build
	@./bin/admin-panel

test:
	@clear
	@go test ./... -v
