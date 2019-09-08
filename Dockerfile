# Stage 1 - Build app
FROM golang:1.12.1-alpine3.9@sha256:e0660b4f1e68e0d408420acb874b396fc6dd25e7c1d03ad36e7d6d1155a4dff6 as build

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

