BINARY_NAME=hb
VERSION=1.0.0
BUILD_DIR=bin

all: clean build

build:
	@echo "🔨 Building hb for current OS..."
	go build -o $(BUILD_DIR)/$(BINARY_NAME) .

build-all: clean
	@echo "🌎 Building hb for multiple platforms..."
	# macOS (M1/M2)
	GOOS=darwin GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 .
	# macOS (Intel)
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 .
	# Linux
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 .
	# Windows
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe .
	@echo "✅ Binaries ready in $(BUILD_DIR)/"

clean:
	@rm -rf $(BUILD_DIR)

test:
	go test ./...

.PHONY: all build build-all clean test
