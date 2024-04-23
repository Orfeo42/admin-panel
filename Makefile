build:
	@templ generate
	@go build -o bin/admin-panel


run: build
	@./bin/admin-panel


# run:
# 	@templ generate
# 	@go run main.go