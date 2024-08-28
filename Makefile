db-up:
	@docker-compose -f docker-compose.yaml up -d postgres
	@echo "🚀 Database is up and running!"

db-down:
	@docker-compose -f docker-compose.yaml down --volumes
	@echo "❄️ Database is down!"

db-init: db-up
	@sleep 5
	@go run ./preload/main.go
	@echo "🤓 Database is initialized!"

db-drop: db-down
	@docker-compose -f docker-compose.yaml rm -v postgres
	@echo "💀 Database deleted!"

db-reset: db-drop db-init

build:
	@templ generate
	@go mod tidy
	@go build -o ./bin/admin-panel ./cmd/api/main.go

run-fresh: build db-reset
	@./bin/admin-panel

run: build db-up
	@./bin/admin-panel

test:
	@clear
	@go test ./... -v
