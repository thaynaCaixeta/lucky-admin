# Configuration
BINARY_NAME=bin/lucky-admin
MAIN_PKG=cmd
SEED_SCRIPT=internal/database/migrations/init.sh

.PHONY: docker-up docker-down docker-logs wire build clean run all \
        seed-local-db reset-local-db run-dev run-dev-all

# Docker commands
docker-up:
	@docker-compose -f docker-compose-local.yml up -d

docker-down:
	@docker-compose -f docker-compose-local.yml down

docker-logs:
	@docker-compose -f docker-compose-local.yml logs -f

# Wire code generation
wire:
	cd $(MAIN_PKG) && wire

# Build the binary (depends on wire)
build: wire
	go build -o $(BINARY_NAME) ./$(MAIN_PKG)

# Run the compiled binary
run: build
	./$(BINARY_NAME)

# Clean binary and generated wire files
clean:
	rm -f $(BINARY_NAME)
	find . -name 'wire_gen.go' -delete

# Default target
run-dev: clean build run

run-dev-all: reset-local-db clean build run

# Full DB Reset + Seed
reset-local-db: docker-down docker-clean docker-up seed-local-db

# Stop and remove Docker containers
docker-down:
	docker-compose -f docker-compose-local.yml down

# Remove persisted DynamoDB volumes
docker-clean:
	rm -rf ./data/dynamodb

# Start up Docker containers with fresh volume
docker-up:
	docker-compose -f docker-compose-local.yml up -d

# Seed DynamoDB local with structure and test data
seed-local-db:
	@echo "Seeding DynamoDB Local..."
	@./$(SEED_SCRIPT)
