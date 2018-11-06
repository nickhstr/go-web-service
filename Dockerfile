# Stage 1 - Build app
FROM golang:1.11.2@sha256:e7462ca504afc789d289f2bb5fd471815cc11833439d2fe4e61915b190045359

# Install git
RUN apk update && apk add git

WORKDIR /app
COPY ./app ./app
COPY ./main.go \
    ./go.mod \
    ./go.sum \
    ./
RUN CGO_ENABLED=0 go build -o ./bin/go-web-service ./main.go

# Stage 2 - Run app
FROM scratch

COPY --from=0 /app/bin/go-web-service go-web-service
ENV GO_ENV=production \
    PORT=3000

EXPOSE 3000
CMD ["./go-web-service"]
