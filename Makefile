include .env
$(eval export $(shell sed -ne 's/ *#.*$$//; /./ s/=.*$$// p' .env))



db-up:
	@mkdir -p /home/${USER}/${APP_NAME}
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
	@sudo rm -r /home/${USER}/${APP_NAME}
	@echo "💀 Database deleted!"

db-reset: db-drop db-init

build:
	@templ generate
	@go mod tidy
	@go build -o ./bin/${APP_NAME} ./cmd/api/main.go

run-fresh: build db-reset
	@./bin/${APP_NAME}

run: build db-up
	@./bin/${APP_NAME}

test:
	@clear
	@go test ./... -v
