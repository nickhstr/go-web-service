# VARIABLES

PROJECTNAME ?= $(shell basename "$(PWD)")
COMMIT := $(shell git rev-parse --short HEAD)
VERSION := $(shell git describe --abbrev=0 --always)
LDFLAGS ?= -ldflags "-X main.gitCommit=$(COMMIT) -X main.appVersion=$(VERSION)"
OUTPUT ?= bin/$(PROJECTNAME)

export GOBIN = $(CURDIR)/bin
export PATH := $(GOBIN):$(PATH)

# TARGETS

# First target, is the default command run if 'make' is invoked without any targets
all: install

## build: Builds the app's executable
.PHONY: build
build:
	@echo "🚧 Building executable..."
	go build -o $(OUTPUT) $(LDFLAGS) $(flags) main.go
	@echo "✨ Done."

## build-prod: Builds the app in a way most suitable for production
.PHONY: build-prod
build-prod:
# Disable cgo to create build that is statically linked
	CGO_ENABLED=0 make build
	upx $(OUTPUT)

## clean: Removes build artifacts
.PHONY: clean
clean:
	@echo "🔥 Removing build artifacts..."
	@rm -rf bin
	@echo "✨ Done."

## coverage: Runs tests and reports coverage
.PHONY: coverage
coverage: create-coverage
	@echo "=============================== Coverage Summary ==============================="
	@go tool cover -func=coverage.out
	@echo "================================================================================"

## coverage-html: Runs tests and opens a browser window to visualize test coverage
.PHONY: coverage-html
coverage-html: create-coverage
	@echo "📊 Opening coverage report in browser..."
	@go tool cover -html=coverage.out

## create-coverage: Outputs test coverage to 'coverage.out'
.PHONY: create-coverage
create-coverage:
	@echo "🏃 Running tests and creating coverage report..."
	@GO_ENV=test go test -race -coverprofile=coverage.out ./...
	@echo "✅ Done."

## dev: Starts the app in dev mode
.PHONY: dev
dev:
	@echo "🚀 Starting dev server..."
	@modd --file=./internal/tools/modd.dev.conf

## install: Downloads/installs all app dependencies
.PHONY: install
install:
	@echo "🚚 Downloading dependencies..."
	@go mod download

	@echo "🛠  Building Go dependencies..."
	@go install github.com/cortesi/modd/cmd/modd
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint
	@go install github.com/psampaz/go-mod-outdated
	@echo "✨ Done."

## lint: Runs linter against Go files
.PHONY: lint
lint:
	@echo "🔍 Linting Go files..."
	@golangci-lint run $(flags)
	@go mod tidy
	@git diff --exit-code -- go.mod go.sum
	@echo "✨ Done."

## mod-outdated: Checks for updates to direct go.mod dependencies
.PHONY: mod-outdated
mod-outdated:
	@go list -u -m -json all | go-mod-outdated -update -direct

## serve: Builds and runs the application in production mode
.PHONY: serve
serve: build
	@echo "🚀 Starting server..."
	@./$(OUTPUT)

## test: Runs all tests
.PHONY: test
test:
	$(eval flags ?= -race)
	$(eval packages ?= ./...)
	@echo "🏃 Running all Go tests..."
	GO_ENV=test go test $(flags) $(packages)
	@echo "✅ Done."

## test-watch: Runs tests and watches for changes
.PHONY: test-watch
test-watch:
	@echo "🏃 Running test watcher..."
	@modd --file=./internal/tools/modd.test.conf

## help: List available commands
.PHONY: help
help: Makefile
	@echo
	@echo " Choose a command to run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/  /'
	@echo
