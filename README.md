# Web Service Boilerplate
Get up and running quickly with a Go web service.

---

## Prerequisites
- Go 1.11 or higher ([Installation Instructions](https://golang.org/doc/install))
- make
- docker ([Installation Instructions](https://www.docker.com/get-started))

## Installation

In a directory outside of the `$GOPATH/src` tree:

```sh
~$ git clone https://github.com/nickhstr/go-web-service.git
~$ cd go-web-service
~$ make install
```

---

## Usage

### Development

For local development, create a `.env` file at the project's root, with the `PORT` variable set to whatever port you desire, otherwise it defaults to 3000. Ex: `PORT=8080`

To start the server:

```sh
~$ make dev
```

That will not only compile and start the server, but it will recompile and restart the app on file changes.

### Production

Ideally, docker would handle production builds.

But to run the app in production mode locally:

```sh
~$ make serve
```

That builds and serves the app. Be sure to specify the `PORT` in the command.

Ex:

```sh
~$ PORT=8080 make serve
```

### Testing

To run all of the app's tests, and print their coverage:

```sh
~$ make test
```

To open the app's test coverage report in a browser:

```sh
~$ make coverage
```
