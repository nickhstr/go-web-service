build:
	go build -o ./bin/go-web-service main.go

clean:
	rm -rf ./bin

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

dev:
	go run main.go

serve: build
	./bin/go-web-service

test:
	go test -v ./...
