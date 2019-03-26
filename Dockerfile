# Stage 1 - Build app
FROM golang:1.12.1-alpine3.9@sha256:5f7781ceb97dd23c28f603c389d71a0ce98f9f6c78aa8cbd12b6ca836bfc6c6c as build

# Install git to fetch dependencies, and install certificates to allow HTTPS requests
RUN apk update && apk add git ca-certificates

WORKDIR /app
COPY . .

RUN go mod download
# Disable cgo to create build that is statically linked
RUN CGO_ENABLED=0 go build -a -o ./bin/service ./main.go

# Stage 2 - Run app
FROM scratch

# Copy the executable
COPY --from=build /app/bin/service service
# Copy certs
COPY --from=build /etc/ssl/ /etc/ssl/

ENV GO_ENV=production \
    PORT=3000

EXPOSE 3000
CMD ["./service"]

