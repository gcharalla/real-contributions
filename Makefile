# Variables
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=real-contributions
MAIN_FILE=cmd/main.go
ENV_FILE=.env
BIN_DIR=bin

# Load environment variables from .env file
include $(ENV_FILE)
export $(shell sed 's/=.*//' $(ENV_FILE))

# Default target
all: build

# Build the binary and move it to the bin directory
build:
	$(GOBUILD) -o $(BIN_DIR)/$(BINARY_NAME) $(MAIN_FILE)

# Clean the project
clean:
	$(GOCLEAN)
	rm -f $(BIN_DIR)/$(BINARY_NAME)

# Run tests
test:
	$(GOTEST) ./...

# Run the application with parameters from .env file
run:
	$(GOBUILD) -o $(BIN_DIR)/$(BINARY_NAME) $(MAIN_FILE)
	./$(BIN_DIR)/$(BINARY_NAME) -add $$ADD_PARAMETER -email $$EMAIL_PARAMETER

# Phony targets
.PHONY: all build clean test deps run
