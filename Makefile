# TIV - Terminal Image Viewer
# Makefile for building and installing

# Variables
BINARY_NAME=tiv
VERSION?=v1.0.0
BUILD_DIR=build
INSTALL_DIR=/usr/local/bin

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

# Linting
GOLANGCI_LINT_VERSION=v1.64.8
GOLANGCI_LINT_CMD=golangci-lint

# Build flags
LDFLAGS=-ldflags "-s -w -X main.Version=$(VERSION)"

.PHONY: all build clean test install uninstall deps tidy lint lint-install check help

# Default target
all: build

# Build the binary
build:
	@echo "Building $(BINARY_NAME) $(VERSION)..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) .

# Build for multiple platforms
build-all:
	@echo "Building for multiple platforms..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 .
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 .
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe .

# Clean build artifacts
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)
	rm -f $(BINARY_NAME)

# Run tests
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

# Install dependencies
deps:
	@echo "Installing dependencies..."
	$(GOMOD) download

# Tidy go.mod
tidy:
	@echo "Tidying go.mod..."
	$(GOMOD) tidy

# Install golangci-lint
lint-install:
	@echo "Installing golangci-lint $(GOLANGCI_LINT_VERSION)..."
	@which $(GOLANGCI_LINT_CMD) > /dev/null 2>&1 || { \
		echo "golangci-lint not found, installing..."; \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin $(GOLANGCI_LINT_VERSION); \
	}
	@echo "golangci-lint installed successfully"

# Run linting (install golangci-lint if needed)
lint: lint-install
	@echo "Running golangci-lint..."
	@$(GOLANGCI_LINT_CMD) run --out-format=colored-line-number --timeout=5m
	@echo "Linting completed successfully!"

# Pre-commit checks (lint + test)
check: lint test
	@echo "All pre-commit checks passed!"

# Install binary to system
install: build
	@echo "Installing $(BINARY_NAME) to $(INSTALL_DIR)..."
	@sudo cp $(BUILD_DIR)/$(BINARY_NAME) $(INSTALL_DIR)/
	@echo "Installation complete. Try: $(BINARY_NAME) --help"

# Uninstall binary from system
uninstall:
	@echo "Uninstalling $(BINARY_NAME)..."
	@sudo rm -f $(INSTALL_DIR)/$(BINARY_NAME)

# Development build (quick)
dev:
	$(GOBUILD) -o $(BINARY_NAME) .

# Release build
release: clean build-all
	@echo "Creating release archives..."
	@cd $(BUILD_DIR) && \
	tar -czf $(BINARY_NAME)-$(VERSION)-linux-amd64.tar.gz $(BINARY_NAME)-linux-amd64 && \
	tar -czf $(BINARY_NAME)-$(VERSION)-darwin-amd64.tar.gz $(BINARY_NAME)-darwin-amd64 && \
	tar -czf $(BINARY_NAME)-$(VERSION)-darwin-arm64.tar.gz $(BINARY_NAME)-darwin-arm64 && \
	zip $(BINARY_NAME)-$(VERSION)-windows-amd64.zip $(BINARY_NAME)-windows-amd64.exe
	@echo "Release archives created in $(BUILD_DIR)/"

# Show help
help:
	@echo "TIV - Terminal Image Viewer"
	@echo ""
	@echo "Available targets:"
	@echo "  build      - Build the binary"
	@echo "  build-all  - Build for multiple platforms"
	@echo "  clean      - Clean build artifacts"
	@echo "  test       - Run tests"
	@echo "  lint       - Run golangci-lint (installs if needed)"
	@echo "  lint-install - Install golangci-lint"
	@echo "  check      - Run lint + test (recommended before commit)"
	@echo "  deps       - Install dependencies"
	@echo "  tidy       - Tidy go.mod"
	@echo "  install    - Install to system (requires sudo)"
	@echo "  uninstall  - Remove from system (requires sudo)"
	@echo "  dev        - Quick development build"
	@echo "  release    - Create release archives"
	@echo "  help       - Show this help" 