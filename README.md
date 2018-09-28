# Web Service Boilerplate
Get up and running quickly with a Go web service.

## Prerequisites
- Go 1.11 or higher ([Installation Instructions](www.google.com))
- modd ([Installation Instructions](https://github.com/cortesi/modd/issues/57))
- make
- docker ([Installation Instructions](https://www.docker.com/get-started))

## Installation

In a directory outside of the `$GOPATH/src` tree:

```sh
~/go-web-service$ git clone https://github.com/nickhstr/go-web-service.git
~/go-web-service$ cd go-web-service
~/go-web-service$ make install
```

## Usage

### Development

For local development, create a `.env` file at the project's root, with the `PORT` variable set to whatever port you desire, otherwise it defaults to 3000. Ex: `PORT=8080`

To start the server:

```sh
~/go-web-service$ make dev
```

That will not only compile and start the server, but it will recompile and restart the app on file changes.

### Production

Ideally, docker would handle production builds.

But to run the app in production mode locally:

```sh
~/go-web-service$ make serve
```

That builds and serves the app. Be sure to specify the `PORT` in the command.

Ex:

```sh
~/go-web-service$ PORT=8080 make serve
```

### Testing

To run all of the app's tests, and print their coverage:

```sh
~/go-web-service$ make test
```

To open the app's test coverage report in a browser:

```sh
~/go-web-service$ make coverage
```

To test individual packages, for example the router's `handlers` package:

```sh
~/go-web-service$ cd app/router/handlers
~/go-web-service/app/router/handlers$ go test
```

or

```sh
~/go-web-service$ go test github.com/nickhstr/go-web-service/app/router/handlers
```
