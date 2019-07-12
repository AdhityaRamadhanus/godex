.PHONY: default test build

OS := $(shell uname)
VERSION ?= 1.0.0

# target #

default: run

run:
	go run cmd/main.go

build: 
	@echo "Setup godex"
ifeq ($(OS),Linux)
	@echo "Build godex..."
	GOOS=linux  go build -ldflags "-s -w -X main.Version=$(VERSION)" -o godex cmd/main.go
endif
ifeq ($(OS) ,Darwin)
	@echo "Build godex..."
	GOOS=darwin go build -ldflags "-X main.Version=$(VERSION)" -o godex cmd/main.go
endif
	@echo "Succesfully Build for ${OS} version:= ${VERSION}"

# Test Packages

test:
	go test -v --cover ./...