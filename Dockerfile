FROM golang:1.13.1-alpine3.10@sha256:2293e952c79b8b3a987e1e09d48b6aa403d703cef9a8fa316d30ba2918d37367 as build

# Install git to fetch dependencies, and install certificates to allow HTTPS requests
RUN apk update && apk add --no-cache git ca-certificates

WORKDIR /app
COPY . .

RUN go get -mod=readonly
# Disable cgo to create build that is statically linked
RUN CGO_ENABLED=0 go build -a -o ./bin/service ./main.go

FROM alpine:3.10

# Copy the executable
COPY --from=build /app/bin/service service
# Copy certs
COPY --from=build /etc/ssl/ /etc/ssl/

ENV GO_ENV=production \
    PORT=3000

EXPOSE 3000
CMD ["./service"]

