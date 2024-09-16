# Define the Go command
GO = go

# Define the target binary name
SVC_NAME = go-jwt-security

# Define the executable binary name
APP_EXECUTABLE = go-jwt-security

# Default target: build the application
all: build

# Build the Go application
build:
	$(GO) build -a -o $(APP_EXECUTABLE) ./cmd/$(SVC_NAME)