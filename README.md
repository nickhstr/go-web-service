# Web Service Boilerplate
Get up and running quickly with a Go web service.

- [Web Service Boilerplate](#web-service-boilerplate)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Usage](#usage)
    - [Development](#development)
    - [Testing](#testing)
    - [Linting](#linting)

---

## Prerequisites
- Go 1.13 or higher ([Installation Instructions](https://golang.org/doc/install))
  - For macOS, `brew` works well
- docker ([Installation Instructions](https://www.docker.com/get-started))

## Installation

```sh
git clone https://github.com/nickhstr/go-web-service.git
cd go-web-service
make
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

### Testing

To run all of the app's tests, and print their coverage:

```sh
make test
```

To run all tests, and watch for changes:
```sh
make test-watch
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
