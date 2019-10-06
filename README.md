# Web Service Boilerplate
Get up and running quickly with a Go web service.

---

## Prerequisites
- Go 1.13 or higher ([Installation Instructions](https://golang.org/doc/install))
- direnv ([Installation Instructions](https://direnv.net/#basic-installation))
- docker ([Installation Instructions](https://www.docker.com/get-started))

## Installation

Note: If using a go version less than 1.13, change to a directory outside of the `$GOPATH/src` tree.

```sh
git clone https://github.com/nickhstr/go-web-service.git
cd go-web-service
make install
```

---

## Usage

### Development

For local development, create a `.env` file at the project's root, with the `PORT` variable set to whatever port you desire, otherwise it defaults to 3000. Ex: `PORT=8080`

To start the server:

```sh
make dev
```

That will not only compile and start the server, but it will recompile and restart the app on file changes.

### Production

Ideally, docker would handle production builds.

But to run the app in production mode locally:

```sh
make serve
```

### Testing

To run all of the app's tests, and print their coverage:

```sh
make test
```

To open the app's test coverage report in a browser:

```sh
make coverage-html
```

### Linting

Lint all `.go` files in the repo:

```sh
make lint
```
