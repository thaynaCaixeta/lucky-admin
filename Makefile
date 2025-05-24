# Configuration
BINARY_NAME=bin/lucky-admin
MAIN_PKG=cmd

.PHONY: docker-up docker-down docker-logs wire build clean run all

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
all: build

run-dev: clean build run