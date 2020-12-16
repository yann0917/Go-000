# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=build/app
BINARY_UNIX=$(BINARY_NAME)_unix

default: run

build:
	$(GOBUILD) $(RACE) -o $(BINARY_NAME) -v ./cmd

test:
	$(GOTEST) $(RACE) -v ./test

enable-race:
	$(eval RACE = -race)

run:
	$(GOBUILD) $(RACE) -o $(BINARY_NAME) -v ./cmd
	cd $(patsubst %/app,%, $(BINARY_NAME)) && ./app

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)


build-race: enable-race build
test-race: enable-race test

.PHONY: build test clean run build-race test-race enable-race