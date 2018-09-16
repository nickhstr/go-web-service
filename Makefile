build:
	go build -o ./bin/go-web-service main.go

clean:
	rm -rf ./bin

coverage:
	go test -cover ./...

dev:
	go run main.go

serve: build
	./bin/go-web-service

test:
	go test -v ./...
