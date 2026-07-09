
.PHONY: help dev build test test-coverage test-integration test-e2e clean migrate-up migrate-down docker-up docker-down lint fmt

#colour output
BLUE := \033[0;34m]
GREEN := \033[0;32m]
RED := \033[0;31m]
YELLOW := \033[1;33m]
NC:=\033[0m]#NO COLOUR


help: #run the help message
	@echo '$(GREEN)Available commands: $(NC)'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST)|sort|awk 'BEGIN {FS =":.*?##"};{printf "$(GREEN)%- 20s $(NC) %s \n", $$1,$$2}'

dev:# run in development mode 
	@echo '$(GREEN)Starting development server...$(NC)'
	go run cmd/server/main.go

build:#builds the application
	@echo '$(GREEN)Building the application...$(NC)'
	CGO_ENABLED =0 GOOS=linux GOARCH =amd64 go build -ldflags="-w -s" -o bin/wallet-server cmd/server/main.go

test:#run all the test
	@echo'$(GREEN)Running all the tests...$(NC)'
	go run -v -race ./...

test-unit:#running unit test only
	@echo '$(GREEN)Running unit tests...$(NC)'
	go run -v -race ./tests/unit/...

test-integration:#Running only integration test
	@echo '$(GREEN)Running integration tests only...$(NC)'
	go run -v -race ./tests/integration/...

test-e2e:#running E2E test onl
	@echo '$(GREEN)Running E2E tests...$(NC)'
	go run -v -race ./tests/e2e/...

test-coverage:#generate test coverage report
	@echo '$(GREEN)Generating coverage report...$(NC)'
	go test -coverprofile = coverage.out ./...
	go tool cover -html = coverage.out -o coverage.html
	@echo '(GREEN)Cover Report Generated: coverage.html $(NC)'

lint: #run linter
	@echo '$(GREEN)Running linter...$(NC)'
	golangci-lint run ./...

fmt:#format code
	@echo '$(GREEN)Formatting code...$(NC)'
	go fmt ./...

clean:# clean build artiufacts
	@echo '$(GREEN)Cleaning...$(NC)'
	rm -rf bin/
	rm -rf tmp/
	rm -rf coverage.out coverage.html


migrate-up:#run database migrations
	@echo '$(GREEN)Running migrations...$(NC)'
	@read -p "Enter database password: " pwd;\
	PGPASSWORD = $$pwd psql -U $(DB_USER) -d $(DB_NAME) -f migrations/001_initial_schema.sql

migrate-down:#run rollback last migration
	@echo '$(GREEN)Rolling back migration...$(NC)'
	@read -p "Enter database password: " pwd;\
	PGPASSWORD = $$pwd psql -U $(DB_USER) -d $(DB_NAME) -f migrations/rollback.sql

docker-up:#start docker container
	@echo '$(GREEN)Starting Docker container...$(NC)'
	docker-compose up -d

docker-down:# stop docker file
	@echo '$(GREEN)Stopping Docker container...$(NC)'
	docker-compose down


.PHONY: help dev build test test-coverage test-integration test-e2e clean migrate-up migrate-down docker-up docker-down lint fmt