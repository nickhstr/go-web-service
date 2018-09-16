# Web Service Boilerplate
This provides boilerplate code for a Go web service.

## Prerequisites
- Go 1.11 or higher ([Installation Instructions](www.google.com))
- make
- docker ([Installation Instructions](https://www.docker.com/get-started))

## Installation

In a directory outside of the `$GOPATH/src` tree:

```sh
$ git clone https://github.com/nickhstr/go-web-service.git
$ cd go-web-service
$ go get
```

## Usage

### Development

For local development, create a `.env` file at the project's root, with the `PORT` variable set to whatever port you desire, otherwise it defaults to 3000.

Ex:

```txt
PORT=8080
```

To start the server:

```sh
$ make dev
```

or

```sh
$ go run main.go
```

### Production

Ideally, docker would handle production builds.

But to run the app in production mode locally:

```sh
$ make serve
```

That builds and serves the app. Be sure to specify the `PORT` in the command.

Ex:

```sh
$ PORT=80 make serve
```
