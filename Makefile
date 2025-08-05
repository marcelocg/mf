# MF TOTP Generator Makefile

# Variables
BINARY_NAME=mf
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME=$(shell date -u +%Y-%m-%dT%H:%M:%SZ)
LDFLAGS=-ldflags "-X main.version=$(VERSION) -X main.buildTime=$(BUILD_TIME)"

# Default target
.PHONY: all
all: build

# Build for current platform
.PHONY: build
build:
	go build $(LDFLAGS) -o $(BINARY_NAME)

# Run tests
.PHONY: test
test:
	go test ./...

# Run tests with coverage
.PHONY: test-coverage
test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Clean build artifacts
.PHONY: clean
clean:
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_NAME)-*
	rm -f *.exe
	rm -f coverage.out coverage.html

# Format code
.PHONY: fmt
fmt:
	go fmt ./...

# Lint code
.PHONY: lint
lint:
	go vet ./...

# Tidy dependencies
.PHONY: tidy
tidy:
	go mod tidy

# Build for all platforms
.PHONY: build-all
build-all: build-linux build-windows build-macos

# Build for Linux
.PHONY: build-linux
build-linux:
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BINARY_NAME)-linux-amd64

# Build for Windows
.PHONY: build-windows
build-windows:
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(BINARY_NAME)-windows-amd64.exe

# Build for macOS
.PHONY: build-macos
build-macos:
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BINARY_NAME)-macos-amd64
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o $(BINARY_NAME)-macos-arm64

# Install locally
.PHONY: install
install: build
	cp $(BINARY_NAME) $(HOME)/.local/bin/

# Development workflow
.PHONY: dev
dev: fmt lint test build

# Create checksums for releases
.PHONY: checksums
checksums:
	@echo "Creating checksums..."
	@sha256sum mf-linux-amd64 > checksums.txt
	@sha256sum mf-windows-amd64.exe >> checksums.txt
	@sha256sum mf-macos-amd64 >> checksums.txt
	@sha256sum mf-macos-arm64 >> checksums.txt
	@echo "Checksums created in checksums.txt"

# Release preparation
.PHONY: release
release: clean fmt lint test-coverage build-all checksums

# Help
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build         - Build for current platform"
	@echo "  build-all     - Build for all supported platforms"
	@echo "  build-linux   - Build for Linux"
	@echo "  build-windows - Build for Windows"
	@echo "  build-macos   - Build for macOS"
	@echo "  test          - Run tests"
	@echo "  test-coverage - Run tests with coverage report"
	@echo "  clean         - Clean build artifacts"
	@echo "  fmt           - Format code"
	@echo "  lint          - Lint code"
	@echo "  tidy          - Tidy dependencies"
	@echo "  install       - Install binary to ~/.local/bin"
	@echo "  dev           - Development workflow (fmt, lint, test, build)"
	@echo "  checksums     - Create SHA256 checksums for binaries"
	@echo "  release       - Release preparation (clean, test, build-all, checksums)"
	@echo "  help          - Show this help"