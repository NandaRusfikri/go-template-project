# Set the compiler to use for building the program
CC=go

GOFLAGS=-ldflags="-s -w"
OUTPUT=backend-api

export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0

# Define the build target
build:
	$(CC) build  -o $(OUTPUT) main.go

# Define the clean target
clean:
	rm -f $(OUTPUT)


swag:
	swag init  -g pkg/swagger.go

