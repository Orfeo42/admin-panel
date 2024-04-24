#Da spostare su docker compose
include .env
$(eval export $(shell sed -ne 's/ *#.*$$//; /./ s/=.*$$// p' .env))

build:
	@templ generate
	@go mod tidy
	@go build -o bin/admin-panel cmd/controllers/main.go

run: build db-up
	@./bin/admin-panel

db-up: ## Start db.
	@mkdir -p /home/${USER}/get-good-db
	@docker-compose -f docker/docker-compose.yaml up -d postgres
	@echo "🚀 Database is up and running!"

db-down: ## Stop db.
	@docker-compose -f docker/docker-compose.yaml down --volumes
	@echo " Database is down!"


db-drop: db-down ## Completly delete the db
	@sudo rm -r /home/${USER}/get-good-db
	@echo " Database deleted!"