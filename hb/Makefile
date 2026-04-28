BINARY_NAME=hb
VERSION=5.0.0
BUILD_DIR=bin

all: build

build:
	go build -o $(BINARY_NAME) main.go

build-all:
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 main.go
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 main.go
	GOOS=darwin GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 main.go
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe main.go

clean:
	rm -f $(BINARY_NAME)
	rm -rf $(BUILD_DIR)

test:
	go test -v ./...

.PHONY: all build build-all clean test
