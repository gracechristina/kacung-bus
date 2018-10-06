BIN_NAME = kacung-bus

all: build run

build: 
	@go build -o bin/${BIN_NAME} ./cmd/${BIN_NAME}

get-deps:
	@dep ensure -v

run:
	@./bin/kacung-bus
