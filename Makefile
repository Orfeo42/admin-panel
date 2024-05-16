include .env
$(eval export $(shell sed -ne 's/ *#.*$$//; /./ s/=.*$$// p' .env))



db-up:
	@mkdir -p /home/${USER}/${PROJECT_NAME}
	@docker-compose -f docker/docker-compose.yaml up -d postgres
	@echo "🚀 Database is up and running!"

db-down:
	@docker-compose -f docker/docker-compose.yaml down --volumes
	@echo "❄️ Database is down!"

db-init: db-up
	@sleep 5
	@go run ./cmd/preload/main.go
	@echo "🤓 Database is initialized!"

db-drop: db-down
	@sudo rm -r /home/${USER}/${PROJECT_NAME}
	@echo "💀 Database deleted!"

db-reset: db-drop db-init

build:
	@templ generate
	@go mod tidy
	@go build -o ./bin/${PROJECT_NAME} ./cmd/controllers/main.go

run-fresh: build db-init
	@clear
	@./bin/${PROJECT_NAME}

run: build db-up
	@clear
	@./bin/${PROJECT_NAME}

run-prove:
	@clear
	@go run ./cmd/prove/main.go

test:
	@clear
	@go test ./... -v
