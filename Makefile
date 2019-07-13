.PHONY: default test run build build-image

export HEAD_COMMIT_SHA1 ?= $(shell git show -q --format=%h)
OS := $(shell uname)
VERSION ?= 1.0.0

build-image: test
	docker build -t godex:$(HEAD_COMMIT_SHA1) .;

run-image:
	docker run -it godex:$(HEAD_COMMIT_SHA1) $(ARG);

# target #

default: run

run:
	go run cmd/main.go

build: test
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