.PHONY: all test build check format clean help vet lint install-tools

MODULE := github.com/ziguayungui/go-uci/dev
GOFILES := $(shell find . -type f -name '*.go' -not -path "./vendor/*")
GOARCHS := amd64 arm64
GOOS := linux

all: test build

help:
	@echo "Available targets:"
	@echo "  help          - Show this help message"
	@echo "  test          - Run tests for current architecture"
	@echo "  test-cross    - Run tests for all supported architectures"
	@echo "  build         - Build for all supported architectures (library, no binary output)"
	@echo "  build-cross   - Build and test for all architectures"
	@echo "  format        - Format code with go fmt"
	@echo "  vet           - Run go vet"
	@echo "  lint          - Run linter"
	@echo "  check         - Run all checks (format, vet, lint, test)"
	@echo "  clean         - Clean build artifacts"
	@echo "  coverage      - Generate coverage report"

test:
	go test -v ./...

test-cross:
	@for arch in $(GOARCHS); do \
		echo "Testing for $(GOOS)/$$arch..."; \
		GOOS=$(GOOS) GOARCH=$$arch CGO_ENABLED=0 go test -v ./...; \
	done

build:
	@for arch in $(GOARCHS); do \
		echo "Building for $(GOOS)/$$arch..."; \
		GOOS=$(GOOS) GOARCH=$$arch CGO_ENABLED=0 go build -v ./...; \
	done

build-cross: build test-cross

format:
	go fmt ./...

vet:
	go vet ./...

lint:
	@if command -v golint >/dev/null 2>&1; then \
		golint ./...; \
	else \
		echo "golint not found. Run 'make install-tools' to install it."; \
	fi

check: format vet test

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

install-tools:
	go install golang.org/x/lint/golint@latest

clean:
	rm -f coverage.out coverage.html
	go clean -cache
	@echo "Clean completed"
