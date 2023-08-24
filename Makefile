# Makefile for building a Go application

# Binary output name
BINARY_NAME=dnsbl-blacklist-checker

# Go build command
build:
	go build -o $(BINARY_NAME)

# Run command
run: build
	./$(BINARY_NAME)

# Clean up command
clean:
	go clean
	rm -f $(BINARY_NAME)

# Build for multiple platforms (Linux, macOS, Windows)
build-all:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)_linux
	GOOS=darwin GOARCH=amd64 go build -o $(BINARY_NAME)_macos
	GOOS=windows GOARCH=amd64 go build -o $(BINARY_NAME)_windows.exe

.PHONY: build run clean build-all