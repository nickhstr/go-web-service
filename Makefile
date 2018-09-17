# The commands listed here merely serve as a convenience.
# Every one of these commands can be ran using the 'go' tool.

PROJECTNAME=$(shell basename "$(PWD)")

# First target, is the default command run if 'make' is invoked without any targets
all: help

## build: Uses 'go build' to create the application executable, found in the 'bin' directory
.PHONY: build
build:
	@echo Building executable...
	@go build -o bin/$(PROJECTNAME) src/main.go
	@echo
	@echo Done

## clean: Removes build artifacts
.PHONY: clean
clean:
	@echo Removing build artifacts...
	@rm -rf bin
	@echo Done

## create-coverage: Outputs test coverage to 'coverage.out'
.PHONY: create-coverage
create-coverage:
	@echo Running tests and creating coverage report...
	@echo Done

## coverage: Runs tests and opens a browser window to visualize test coverage
.PHONY: coverage
coverage: create-coverage
	@echo Opening coverage report in browser...
	@go tool cover -html=coverage.out
	@echo Done

## dev: Starts the app in dev mode
.PHONY: dev
dev:
	@echo Starting dev server...
	@go run src/main.go

## install: downloads all app dependencies
.PHONY: install
install:
	@echo Installing package dependencies...
	@go get ./...
	@echo Done

## serve: Builds and runs the application in production mode
.PHONY: serve
serve: build
	GO_ENV=production ./bin/$(PROJECTNAME)

## test: Runs all tests
.PHONY: test
test:
	@echo Running all application tests...
	@go test -v ./...
	@echo Done

print:
	@echo $(PROJECTNAME)

## help: List available commands
.PHONY: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
