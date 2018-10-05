# The commands listed here merely serve as a convenience.
# Every one of these commands can be ran using the 'go' tool.

PROJECTNAME=$(shell basename "$(PWD)")
APP_DEV_NAME="app-dev"

# First target, is the default command run if 'make' is invoked without any targets
all: help

## build: Uses 'go build' to create the application executable, found in the 'bin' directory
.PHONY: build
build:
	@echo "> Building executable..."
	@go build -o bin/$(PROJECTNAME) main.go

## clean: Removes build artifacts
.PHONY: clean
clean:
	@echo "> Removing build artifacts..."
	@rm -rf bin

## create-coverage: Outputs test coverage to 'coverage.out'
.PHONY: create-coverage
create-coverage:
	@echo "> Running tests and creating coverage report..."
	@go test -coverprofile=coverage.out ./...

## coverage: Runs tests and opens a browser window to visualize test coverage
.PHONY: coverage
coverage: create-coverage
	@echo "> Opening coverage report in browser..."
	@go tool cover -html=coverage.out

## dev: Starts the app in dev mode
.PHONY: dev
dev:
	@echo "> Starting dev server..."
	@modd

## install: Downloads all app dependencies
.PHONY: install
install:
	@echo "> Installing package dependencies..."
	@go get ./...
	@go mod tidy

## serve: Builds and runs the application in production mode
.PHONY: build
serve: build
	GO_ENV=production ./bin/$(PROJECTNAME)

## test: Runs all tests
.PHONY: test
test:
	@echo "> Running all application tests..."
	@go test -coverprofile=coverage.out ./...
	@echo Coverage:
	@go tool cover -func=coverage.out

## help: List available commands
.PHONY: help
help: Makefile
	@echo
	@echo " Choose a command to run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
