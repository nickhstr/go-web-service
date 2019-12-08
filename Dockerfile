FROM golang:1.13.4-alpine3.10@sha256:9d2a7c5b6447f525da0a4f18efd2cb05bf7d70228f75d713b7a67345f30157ac as build

# Install dependencies:
# - git to fetch Go dependencies
# - make to run predefined scripts
# - ca-certificates to allow HTTPS requests
# - upx to compress executable
RUN apk update && apk add --no-cache \
  git \
  make \
  ca-certificates \
  upx

WORKDIR /app

# Copy over files needed for `go mod download`
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy all source files
COPY . .

RUN make build-prod OUTPUT=bin/service


FROM alpine:3.10@sha256:c19173c5ada610a5989151111163d28a67368362762534d8a8121ce95cf2bd5a

# Copy the executable
COPY --from=build /app/bin/service service
# Copy certs
COPY --from=build /etc/ssl/ /etc/ssl/

ENV GO_ENV=production \
  PORT=3000

EXPOSE $PORT
CMD ["./service"]
