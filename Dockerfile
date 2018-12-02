# Stage 1 - Build app
FROM golang:1.11.2-alpine3.8@sha256:e7462ca504afc789d289f2bb5fd471815cc11833439d2fe4e61915b190045359 as build

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

