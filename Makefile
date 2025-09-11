# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

# ==================================================================================== #
# VARIABLES
# ==================================================================================== #

# Database
DB_HOST=
DB_USER=
DB_PASSWORD=
DB_NAME=
DB_PORT=5432
DB_SSLMODE=disable
DB_DSN=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)
MIGRATIONS_DIR=sql/schema

# Application
BINARY_NAME=
BUILD_DIR=./bin

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## run: run the application
.PHONY: run
run:
	go run ./cmd/api

## build: build the application
.PHONY: build
build:
	@echo "Building..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/api

## clean: clean build directory
.PHONY: clean
clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)

## test: run all tests
.PHONY: test
test:
	go test -v ./...

## test/cover: run all tests with coverage
.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

# ==================================================================================== #
# DATABASE
# ==================================================================================== #

## db/status: check migration status
.PHONY: db/status
db/status:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_DSN)" status

## db/up: run all pending migrations
.PHONY: db/up
db/up:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_DSN)" up

## db/down: rollback last migration
.PHONY: db/down
db/down:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_DSN)" down

## db/reset: reset database (down all + up all)
.PHONY: db/reset
db/reset:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_DSN)" reset

## db/create: create a new migration file
.PHONY: db/create
db/create:
	@mkdir -p $(MIGRATIONS_DIR)
	@read -p "Enter migration name: " name; \
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_DSN)" create $name sql

## db/backup: create database backup
.PHONY: db/backup
db/backup:
	@echo "Creating backup..."
	@mkdir -p ./backups
	pg_dump "$(DB_DSN)" > ./backups/backup_$(shell date +%Y%m%d_%H%M%S).sql
	@echo "Backup created in ./backups/"

## db/restore: restore database from backup (usage: make db/restore FILE=backup.sql)
.PHONY: db/restore
db/restore:
	@if [ -z "$(FILE)" ]; then echo "Usage: make db/restore FILE=backup.sql"; exit 1; fi
	psql "$(DB_DSN)" < $(FILE)

## db/connect: connect to database
.PHONY: db/connect
db/connect:
	psql "$(DB_DSN)"

## db/drop: drop database (DANGEROUS!)
.PHONY: db/drop
db/drop:
	@echo "Are you sure you want to drop the database? This cannot be undone!"
	@read -p "Type 'yes' to confirm: " confirm; \
	if [ "$$confirm" = "yes" ]; then \
		dropdb -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) $(DB_NAME); \
		echo "Database dropped!"; \
	else \
		echo "Cancelled."; \
	fi

## db/create-db: create database
.PHONY: db/create-db
db/create-db:
	createdb -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) $(DB_NAME)

# ==================================================================================== #
# DEPENDENCIES
# ==================================================================================== #

## deps: install dependencies
.PHONY: deps
deps:
	go mod download
	go mod verify

## deps/update: update dependencies
.PHONY: deps/update
deps/update:
	go get -u ./...
	go mod tidy

## deps/vendor: create vendor directory
.PHONY: deps/vendor
deps/vendor:
	go mod vendor

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## audit: tidy dependencies and format, vet and test all code
.PHONY: audit
audit:
	@echo "Tidying and verifying module dependencies..."
	go mod tidy
	go mod verify
	@echo "Formatting code..."
	go fmt ./...
	@echo "Vetting code..."
	go vet ./...
	staticcheck ./...
	@echo "Running tests..."
	go test -race -vet=off ./...

## fmt: format all go files
.PHONY: fmt
fmt:
	go fmt ./...

## vet: vet all go files
.PHONY: vet
vet:
	go vet ./...

## lint: run golangci-lint
.PHONY: lint
lint:
	golangci-lint run

# ==================================================================================== #
# PRODUCTION
# ==================================================================================== #

## production/deploy: deploy to production (with backup)
.PHONY: production/deploy
production/deploy:
	@echo "🚀 Starting production deployment..."
	@$(MAKE) db/backup
	@echo "🔄 Running migrations..."
	@$(MAKE) db/up
	@echo "🏗️  Building application..."
	@$(MAKE) build
	@echo "✅ Deployment complete!"

## production/backup: create production backup
.PHONY: production/backup
production/backup: db/backup

# ==================================================================================== #
# DOCKER (Optional)
# ==================================================================================== #

## docker/build: build docker image
.PHONY: docker/build
docker/build:
	docker build -t $(BINARY_NAME) .

## docker/run: run docker container
.PHONY: docker/run
docker/run:
	docker run --rm -p 8080:8080 $(BINARY_NAME)

# ==================================================================================== #
# INSTALLATION
# ==================================================================================== #

## install/goose: install goose migration tool
.PHONY: install/goose
install/goose:
	go install github.com/pressly/goose/v3/cmd/goose@latest

## install/tools: install all development tools
.PHONY: install/tools
install/tools:
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest