.PHONY: build
build:
	go build -o bin/go-web-service src/main.go

.PHONY: clean
clean:
	rm -rf bin

.PHONY: create-coverage
create-coverage:
	go test -coverprofile=coverage.out ./...

.PHONY: coverage
coverage: create-coverage
	go tool cover -html=coverage.out

.PHONY: dev
dev:
	go run src/main.go

.PHONY: serve
serve: build
	GO_ENV=production ./bin/go-web-service

.PHONY: test
test:
	go test -v ./...
