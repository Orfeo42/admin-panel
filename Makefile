build:
	@templ generate
	@go mod tidy
	@go build -o bin/admin-panel cmd/controllers/main.go


run: build
	@./bin/admin-panel
