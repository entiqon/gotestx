# GoTestX Makefile

BINARY := gotestx
VERSION := $(shell git describe --tags --abbrev=0)
COMMIT := $(shell git rev-parse --short HEAD)

GO := go
BUILD_DIR := bin

.PHONY: build test install clean release help

help:
	@echo "GoTestX build commands"
	@echo ""
	@echo "make build      Build local binary"
	@echo "make test       Run tests"
	@echo "make install    Install tool locally"
	@echo "make release    Build binaries for multiple platforms"
	@echo "make clean      Remove build artifacts"

build:
	@echo "Building $(BINARY)..."
	$(GO) build -o $(BUILD_DIR)/$(BINARY) .

test:
	@echo "Running tests..."
	$(GO) test ./internal/... -v

install:
	@echo "Installing $(BINARY)..."
	$(GO) install .

clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(BUILD_DIR)

release: clean
	@echo "Building release binaries..."
	mkdir -p $(BUILD_DIR)

	GOOS=darwin GOARCH=arm64 $(GO) build -o $(BUILD_DIR)/$(BINARY)-darwin-arm64 .
	GOOS=darwin GOARCH=amd64 $(GO) build -o $(BUILD_DIR)/$(BINARY)-darwin-amd64 .
	GOOS=linux GOARCH=amd64 $(GO) build -o $(BUILD_DIR)/$(BINARY)-linux-amd64 .
	GOOS=windows GOARCH=amd64 $(GO) build -o $(BUILD_DIR)/$(BINARY)-windows-amd64.exe .

	@echo ""
	@echo "Release binaries created in ./bin"